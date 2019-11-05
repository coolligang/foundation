package httpexpect

import (
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"testing"
	"time"
)

func NewHttpExpect(t *testing.T, baseURL string) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		// prepend this url to all requests
		BaseURL: baseURL,

		// use http.Client with a cookie jar and timeout
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: time.Second * 30,
		},

		// use fatal failures
		Reporter: httpexpect.NewRequireReporter(t),

		// use verbose logging
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func NewHttpExpectWithTimeout(t *testing.T, baseURL string, timeout time.Duration) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		// prepend this url to all requests
		BaseURL: baseURL,

		// use http.Client with a cookie jar and timeout
		Client: &http.Client{
			Jar:     httpexpect.NewJar(),
			Timeout: timeout * time.Second,
		},

		// use fatal failures
		Reporter: httpexpect.NewRequireReporter(t),

		// use verbose logging
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}
