package yandexDiskApiClient

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var logger = os.Stdout

type loggingRoundTripper struct {
	next http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

var transport = loggingRoundTripper{
	next: http.DefaultTransport,
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	fmt.Fprintf(logger, "REDIRECT: %s", req.Response.Status)

	return nil
}
