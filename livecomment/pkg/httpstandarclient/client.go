package httpstandarclient

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func DefaultHTTPConfig(retryMax int, timeOutSec int) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = retryMax
	httpClient := retryClient.StandardClient()
	httpClient.Timeout = time.Duration(timeOutSec) * time.Second
	return httpClient
}
