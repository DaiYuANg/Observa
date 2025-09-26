package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
)

func main() {
	opts := &server.Options{}
	ns, err := server.NewServer(opts)

	if err != nil {
		fmt.Errorf("server start failed: %v", err)
		panic(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}
	ns.WaitForShutdown()
}
