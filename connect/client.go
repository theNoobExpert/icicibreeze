package connect

import (
	"github.com/theNoobExpert/icicibreeze/pkg/transports"
	"net/http"
	"time"
)

func NewHttpClient(timeout time.Duration) *http.Client {
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &http.Client{
		Timeout: timeout,
		Transport: &client.LoggingTransport{
			Transport: &client.HeaderTransport{},
		},
	}
}
