package main

import (
	"fmt"
	"github.com/xeb/backq/modules/certgen"
	"github.com/xeb/backq/modules/public"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

var (
	reqport  = kingpin.Flag("request_port", "The 0MQ port for publishing requests to bqprivate, e.g. a value of 20000 means binding to 'tcp://*:20000'").Required().Int()
	repport  = kingpin.Flag("reply_port", "The 0MQ port for listening for replies from bqprivate").Required().Int()
	httpport = kingpin.Flag("http_port", "The HTTP Port to listen on").Required().Int()
	certpath = kingpin.Flag("certpath", "If set, this will Enable SSL and set the Cert Path").String()
	keypath  = kingpin.Flag("keypath", "If set, this will Enable SSL and set the Key Path").String()
	genssl   = kingpin.Flag("gen_ssl", "If set, this will automatically create cert.pem and key.pem and enable SSL").Bool()
)

func main() {
	kingpin.Parse()

	fmt.Printf("[PUBLIC] Using default port %d\n", *httpport)

	reqaddy := fmt.Sprintf("tcp://*:%d", *reqport)
	repaddy := fmt.Sprintf("tcp://*:%d", *repport)

	fmt.Printf("[PUBLIC] Binding request-0mq channel to '%s'\n", reqaddy)
	fmt.Printf("[PUBLIC] Binding reply-0mq channel to '%s'\n", repaddy)
	fmt.Printf("[PUBLIC] Binding HTTP receiver to ':%d'\n", *httpport)

	if *genssl {
		certgen.Generate("localhost", "Jan 1 15:04:05 2011", 365*24*time.Hour, false, 2048, "P256")
		*certpath = "cert.pem"
		*keypath = "key.pem"
	}

	public.BindBackQ(reqaddy, repaddy)
	public.BindHTTP(*httpport, *certpath, *keypath)
}
