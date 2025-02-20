package client

import (
	"fmt"
	"net/http"
	"time"
)

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
		fmt.Println("Retrying...", i+1)
		time.Sleep(time.Second)
	}

	return resp, err
}

type LoggingTransport struct {
	Transport http.RoundTripper
}

func (lt *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	fmt.Println("Request:", req.Method, req.URL)

	fmt.Println("Request body: ", req.Body)

	if lt.Transport == nil {
		lt.Transport = http.DefaultTransport
	}

	resp, err := lt.Transport.RoundTrip(req)

	if err == nil {
		fmt.Println("Response Status:", resp.Status, "Time:", time.Since(start))
	}

	return resp, err
}
