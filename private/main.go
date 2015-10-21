package main

import (
	"fmt"

	"github.com/xeb/backq/modules/private"
)

const reqaddy string = "tcp://localhost:20000"
const repaddy string = "tcp://localhost:30000"

func main() {
	fmt.Printf("[PRIVATE] Binding request-0mq channel to '%s'\n", reqaddy)
	fmt.Printf("[PRIVATE] Binding reply-0mq channel to '%s'\n", repaddy)

	private.Subscribe(reqaddy, repaddy)
}
