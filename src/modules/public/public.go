package public

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/xeb/backq/src/modules/messages"
)

var reqsock *zmq.Socket
var repsock *zmq.Socket

func BindBackQ(reqaddy string, repaddy string) {
	reqsock, _ = zmq.NewSocket(zmq.PUB)
	reqsock.Bind(reqaddy)

	repsock, _ = zmq.NewSocket(zmq.REP)
	repsock.Bind(repaddy)
}

func sendRequest(r *messages.Request) (rep *messages.Reply) {

	rep = &messages.Reply{}

	v, e := reqsock.Send(r.String(), 0)
	fmt.Printf("[PUBLIC] dispatched request of %s with %d bytes and error is: %s\n", r.String(), v, e)

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
		fmt.Printf("[PUBLIC] *received* %s\n", reply)
		rep.StatusCode = 200
	case <-time.After(time.Second * 1):
		fmt.Printf("[PUBLIC] TIMEOUT waiting for reply\n")
		rep.StatusCode = 500
	}

	return rep
}

func BindHttp(port int) {
	http.HandleFunc("/", handleRequest)
	address := fmt.Sprintf(":%d", port)
	println("[PUBLIC] Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("ListenAndServe ERR %s\n", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %d %s %s\n", 200, r.Method, r.URL.Path)

	bodybytes, _ := ioutil.ReadAll(r.Body)
	body := string(bodybytes)
	req := &messages.Request{Url: r.URL.String(), Headers: &r.Header, Body: body}

	rep := sendRequest(req)

	w.WriteHeader(rep.StatusCode)
	fmt.Fprintf(w, rep.String())
}
