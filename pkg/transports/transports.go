package client

import (
	"net/http"
	"time"

	"github.com/theNoobExpert/icicibreeze/pkg/utils"
)

var logger = utils.GetLogger()

type HeaderTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

func (ht *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range ht.Headers {
		req.Header.Set(key, value)
	}

	if ht.Transport == nil {
		ht.Transport = http.DefaultTransport
	}

	return ht.Transport.RoundTrip(req)
}

type RetryTransport struct {
	Transport  http.RoundTripper
	MaxRetries int
}

func (rt *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.Transport == nil {
		rt.Transport = http.DefaultTransport
	}

	var resp *http.Response
	var err error

	for i := 0; i < rt.MaxRetries; i++ {
		resp, err = rt.Transport.RoundTrip(req)
		if err == nil {
			return resp, nil
		}
		logger.Warnf("Retrying request %s %s (attempt %d/%d): %v", req.Method, req.URL, i+1, rt.MaxRetries, err)
		time.Sleep(time.Second)
	}

	logger.Errorf("Request failed after %d retries: %v", rt.MaxRetries, err)
	return resp, err
}

type LoggingTransport struct {
	Transport http.RoundTripper
}

func (lt *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	logger.Infof("Request: %s %s", req.Method, req.URL)

	if req.Body != nil {
		logger.Debug("Request body: %v", req.Body)
	}

	if lt.Transport == nil {
		lt.Transport = http.DefaultTransport
	}

	resp, err := lt.Transport.RoundTrip(req)

	if err != nil {
		logger.Errorf("Request failed: %s %s, error: %v", req.Method, req.URL, err)
	} else {
		logger.Infof("Response Status: %s, Time Taken: %v", resp.Status, time.Since(start))
	}

	return resp, err
}
