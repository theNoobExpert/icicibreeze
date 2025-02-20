package connect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/theNoobExpert/icicibreeze/pkg/config"
	"github.com/theNoobExpert/icicibreeze/pkg/utils"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type BreezeCredentials struct {
}

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

func (bc *BreezeConnect) MakeRequest(request *BreezeRequest) ([]byte, error) {

	err := utils.Validate.Struct(request)
	if err != nil {
		return nil, fmt.Errorf("error validating request: ")
	}

	var req *http.Request
	var newReqErr error

	if request.Body == "" {
		req, newReqErr = http.NewRequest(string(request.Method), request.URL, nil)
	} else {
		req, newReqErr = http.NewRequest(string(request.Method), request.URL, strings.NewReader(request.Body))
	}

	if newReqErr != nil {
		return nil, fmt.Errorf("could not create request: %w", newReqErr)
	}

	for headerKey, headerValue := range request.Headers {
		req.Header[headerKey] = []string{headerValue} // using req.Header.Set changes the case of header key causing error as headers are case sensitve
	}

	resp, err := bc.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (bc *BreezeConnect) MakeRequestWithTokens(method config.APIRequestMethod, endpoint config.APIEndpoint, payload any, headers map[string]string) ([]byte, error) {
	url := config.API_URL + string(endpoint)
	body := "{}"

	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = string(bodyBytes)
	}

	requestHeaders := bc.GenerateHeaders(body, "")

	if headers != nil {
		for headerKey, headerValue := range headers {
			requestHeaders[headerKey] = headerValue
		}
	}

	return bc.MakeRequest(
		&BreezeRequest{
			Method:  method,
			URL:     url,
			Body:    body,
			Headers: requestHeaders,
		},
	)
}
