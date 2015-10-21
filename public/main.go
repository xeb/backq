package main

import (
	"fmt"

	"github.com/xeb/backq/modules/public"
)

const reqaddy string = "tcp://*:20000"
const repaddy string = "tcp://*:30000"
const httpport int = 9099

func main() {
	fmt.Printf("[PUBLIC] Binding request-0mq channel to '%s'\n", reqaddy)
	fmt.Printf("[PUBLIC] Binding reply-0mq channel to '%s'\n", repaddy)
	fmt.Printf("[PUBLIC] Binding HTTP receiver to ':%d'\n", httpport)

	public.BindBackQ(reqaddy, repaddy)
	public.BindHTTP(httpport)
}
