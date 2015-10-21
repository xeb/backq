package private

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	zmq "github.com/pebbe/zmq4"
	"github.com/xeb/backq/modules/messages"
)

var reqsock *zmq.Socket
var repsock *zmq.Socket
var subscribed = true

// Subscribe will automatically wire up processing of messages for a request address
func Subscribe(reqaddy string, repaddy string) {
	reqsock, _ = zmq.NewSocket(zmq.SUB)
	defer reqsock.Close()
	reqsock.Connect(reqaddy)
	reqsock.SetSubscribe("")

	repsock, _ = zmq.NewSocket(zmq.REQ)
	repsock.Connect(repaddy)
	defer repsock.Close()

	fmt.Println("[PRIVATE] Connected to PUBLIC")

	for subscribed {
		fmt.Println("[PRIVATE] Receiving")
		r, _ := reqsock.Recv(0)
		fmt.Printf("[PRIVATE] Received %d\n", len(r))

		fmt.Printf("[PRIVATE] Executing inner HTTP request\n")

		req, e := messages.NewRequest(r)

		if e != nil {
			fmt.Printf("Terrible error trying to recreate request")
			panic(e)
		}
		reply, body, e := MakeRequest(req)
		SendReply(reply, body, repsock)
	}

	fmt.Println("[PRIVATE] Unsubscribing")
}

// SendReply will send the reply back to the connected reply proxy
func SendReply(r *http.Response, body string, repsock *zmq.Socket) {
	rep := &messages.Reply{StatusCode: r.StatusCode, Body: body, Headers: r.Header}
	reps := rep.String()
	fmt.Printf("[PRIVATE] sending %d bytes or \n----\n%s\n-----\n", len(reps), reps)
	repsock.Send(reps, 0)
	fmt.Printf("[PRIVATE] sent %d bytes\n", len(reps))
}

// MakeRequest will actually go out and make a private HTTP request
func MakeRequest(req *messages.Request) (r *http.Response, body string, e error) {
	client := &http.Client{}
	urlv, _ := url.Parse(fmt.Sprintf("http://www.google.com%s", req.URL))
	request := &http.Request{
		Method: req.Method,
		URL:    urlv,
		Header: req.Headers,
	}
	r, e = client.Do(request)
	defer r.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	body = string(bodyBytes)
	return
}
