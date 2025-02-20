package connect

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/theNoobExpert/icicibreeze/pkg/config"
	"github.com/theNoobExpert/icicibreeze/pkg/utils"
)

var logger = utils.GetLogger()

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type BreezeCredentials struct{}

type BreezeConnect struct {
	Client              HttpClient
	AppKey              string
	AppSecret           string
	ApiSessionKey       string
	ApiSessionToken     string
	IsClientInitialized bool
}

type BreezeRequest struct {
	Method  config.APIRequestMethod `validate:"required"`
	URL     string                  `validate:"required"`
	Body    string
	Headers map[string]string
}

func NewBreezeConnectClient(appKey, appSecret, apiSessionKey string, timeoutSeconds int) (*BreezeConnect, error) {
	httpClient := NewHttpClient(time.Duration(timeoutSeconds) * time.Second)

	breezeClient := &BreezeConnect{
		Client:        httpClient,
		AppKey:        appKey,
		AppSecret:     appSecret,
		ApiSessionKey: apiSessionKey,
	}

	if apiSessionKey != "" {
		_, err := breezeClient.InitSessionToken()
		if err != nil {
			return nil, fmt.Errorf("error while initializing breeze client: %w", err)
		}
	}

	logger.Infof("BreezeConnect client initialized with AppKey: %s", appKey)

	return breezeClient, nil
}

func (brc *BreezeConnect) GenerateHeaders(body string, contentType string) map[string]string {
	if contentType == "" {
		contentType = "application/json"
	}
	checksum, timestamp := utils.CalculateChecksum(body, brc.AppSecret)

	return map[string]string{
		"Connection":     "keep-alive",
		"X-Checksum":     "token " + checksum,
		"X-Timestamp":    timestamp,
		"Content-Type":   contentType,
		"X-AppKey":       brc.AppKey,
		"X-SessionToken": brc.ApiSessionToken,
	}
}

func (brc *BreezeConnect) MakeRequest(request *BreezeRequest) ([]byte, error) {
	logger.Infof("Making request: %s %s", request.Method, request.URL)

	err := utils.Validate.Struct(request)
	if err != nil {
		logger.Errorf("Error validating request: %v", err)
		return nil, err
	}

	var req *http.Request
	var newReqErr error

	if request.Body == "" {
		req, newReqErr = http.NewRequest(string(request.Method), request.URL, nil)
	} else {
		req, newReqErr = http.NewRequest(string(request.Method), request.URL, strings.NewReader(request.Body))
	}

	if newReqErr != nil {
		logger.Errorf("Could not create request: %v", newReqErr)
		return nil, newReqErr
	}

	for headerKey, headerValue := range request.Headers {
		req.Header[headerKey] = []string{headerValue}
	}

	resp, err := brc.Client.Do(req)
	if err != nil {
		logger.Errorf("Failed to execute request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	logger.Infof("Request completed with status: %s", resp.Status)
	return io.ReadAll(resp.Body)
}

func (brc *BreezeConnect) MakeRequestWithTokens(method config.APIRequestMethod, endpoint config.APIEndpoint, payload any, headers map[string]string) ([]byte, error) {

	if brc.ApiSessionToken == "" {
		return nil, errors.New("Api session key not initialized.")
	}

	url := config.API_URL + string(endpoint)
	body := "{}"

	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			logger.Errorf("Failed to marshal request body: %v", err)
			return nil, err
		}
		body = string(bodyBytes)
	}

	requestHeaders := brc.GenerateHeaders(body, "")

	if headers != nil {
		for headerKey, headerValue := range headers {
			requestHeaders[headerKey] = headerValue
		}
	}

	logger.Infof("Making authenticated request to %s", url)
	return brc.MakeRequest(
		&BreezeRequest{
			Method:  method,
			URL:     url,
			Body:    body,
			Headers: requestHeaders,
		},
	)
}
