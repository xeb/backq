package messages

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSerializeRequest(t *testing.T) {
	req := &Request{
		URL:     "http://google.com",
		Headers: make(http.Header),
		Body:    "None",
		Method:  "POST",
	}
	req.Headers["Host"] = []string{"test"}

	result := req.String()
	t.Log(result)

	nreq, e := NewRequest(result)
	if e != nil {
		t.Error(e)
	}

	if nreq.URL != req.URL {
		t.Error(fmt.Printf("URLs do not match, %s vs. %s", nreq.URL, req.URL))
	}
}

func TestSerializeReply(t *testing.T) {
	rep := &Reply{
		StatusCode: 200,
		Headers:    make(http.Header),
		Body:       "None",
	}
	rep.Headers["Host"] = []string{"test"}

	result := rep.String()
	t.Log(result)

	nrep, e := NewReply(result)
	if e != nil {
		t.Error(e)
	}

	if nrep.Body != rep.Body {
		t.Error(fmt.Printf("Bodies do not match, %s vs. %s", nrep.Body, rep.Body))
	}
}

func TestDeserializeRequest(t *testing.T) {
	expected := "{\"URL\":\"/fsdalj\",\"Headers\":{\"Accept\":[\"*/*\"],\"Accepts\":[\"fACE\"],\"User-Agent\":[\"curl/7.43.0\"]},\"Body\":\"\",\"Method\":\"GET\"}"
	nreq, e := NewRequest(expected)
	if e != nil {
		panic(e)
	}

	if nreq.Headers["Accepts"][0] != "fACE" {
		t.Error("No face")
	}
}
