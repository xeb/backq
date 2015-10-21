package public

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/xeb/backq/modules/messages"
)

var reqsock *zmq.Socket
var repsock *zmq.Socket

// BindBackQ will bind the 0MQ backend endpoints specifid
func BindBackQ(reqaddy string, repaddy string) {
	reqsock, _ = zmq.NewSocket(zmq.PUB)
	reqsock.Bind(reqaddy)

	repsock, _ = zmq.NewSocket(zmq.REP)
	repsock.Bind(repaddy)
	// repsock.SetSubscribe("")
}

func sendRequest(r *messages.Request) (rep *messages.Reply) {

	rep = &messages.Reply{}

	v, e := reqsock.Send(r.String(), 0) // send the Request message as a string

	fmt.Printf("[PUBLIC] dispatched request with %d bytes and error is: %s\n", v, e)

	// wait for a reply for a certain period of time
	fmt.Printf("[PUBLIC] Waiting for response...\n")

	var reply string
	rc := make(chan string, 1)
	go func() {
		reply, _ = repsock.Recv(0)
		rc <- reply
	}()

	select {
	case reply = <-rc:
		fmt.Printf("[PUBLIC] received %d bytes \n", len(reply))
		rep.StatusCode = 200
	case <-time.After(time.Second * 10):
		fmt.Printf("[PUBLIC] TIMEOUT waiting for reply\n")
		rep.StatusCode = 500
	}

	return rep
}

// BindHTTP will bind the listener HTTP interface for accepting backq'd requests
func BindHTTP(port int) {
	http.HandleFunc("/", handleRequest)
	address := fmt.Sprintf(":%d", port)
	println("[PUBLIC] Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("ListenAndServe ERR %s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	bodybytes, _ := ioutil.ReadAll(r.Body)
	body := string(bodybytes)
	req := &messages.Request{
		URL:     r.URL.String(),
		Headers: r.Header,
		Body:    body,
		Method:  r.Method,
	}

	rep := sendRequest(req)
	fmt.Printf("[PUBLIC] Writing response %s\n", rep)

	w.WriteHeader(rep.StatusCode)
	for k, _ := range rep.Headers {
		w.Header().Set(k, rep.Headers[k][0])
		fmt.Printf("[PUBLIC] Writing header %s==%s", k, rep.Headers[k][0])
	}

	fmt.Fprintf(w, rep.Body)
	fmt.Printf("[PUBLIC] DONE, request handled %s\n", rep)
}
