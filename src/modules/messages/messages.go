package messages

import (
	"fmt"
	"net/http"
)

type Request struct {
	Url     string
	Headers *http.Header
	Body    string
}

type Reply struct {
	StatusCode int
	Headers    *http.Header
	Body       string
}

func (r *Reply) String() string {
	return fmt.Sprintf("StatusCode=%d,Headers=%s,Body=%s", r.StatusCode, r.Headers, r.Body)
}

func (r *Request) String() string {
	return fmt.Sprintf("Url=%s,Headers=%s,Body=%s", r.Url, r.Headers, r.Body)
}
