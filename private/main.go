package main

import (
	"fmt"
	"github.com/xeb/backq/modules/private"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	reqport    = kingpin.Flag("request_port", "The 0MQ port for publishing requests to bqprivate, e.g. a value of 20000 means binding to 'tcp://*:20000'").Required().Int()
	repport    = kingpin.Flag("reply_port", "The 0MQ port for listening for replies from bqprivate").Required().Int()
	publichost = kingpin.Flag("public_host", "The host name or IP address of the bqpublic server").Required().String()
)

func main() {
	kingpin.Parse()

	reqaddy := fmt.Sprintf("tcp://%s:%d", *publichost, *reqport)
	repaddy := fmt.Sprintf("tcp://%s:%d", *publichost, *repport)

	fmt.Printf("[PRIVATE] Binding request-0mq channel to '%s'\n", reqaddy)
	fmt.Printf("[PRIVATE] Binding reply-0mq channel to '%s'\n", repaddy)

	private.Subscribe(reqaddy, repaddy)
}
