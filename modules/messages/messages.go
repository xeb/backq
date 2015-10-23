package messages

import (
	"encoding/json"
	"net/http"
)

// Request Message representing an HTTP Request
type Request struct {
	URL     string
	Headers http.Header
	Body    string
	Method  string

	// The original value from http.Request.Host
	Host string
}

// Reply Message representing an HTTP Reply
type Reply struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

// Creates a New Request
func NewRequest(s string) (r *Request, e error) {
	var req Request
	e = json.Unmarshal([]byte(s), &req)
	return &req, e
}

func NewReply(s string) (r *Reply, e error) {
	var rep Reply
	e = json.Unmarshal([]byte(s), &rep)
	return &rep, e
}

func (r *Reply) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}

func (r *Request) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}
