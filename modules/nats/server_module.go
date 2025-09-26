package nats

import (
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"go.uber.org/fx"
)

type ReadyChan chan struct{}

var ServerModule = fx.Module("nats_server_module",
	fx.Provide(newServer, newReadyChan),
	fx.Invoke(lifecycle),
)

func newReadyChan() ReadyChan {
	return make(chan struct{}, 1)
}
func newServer() (*server.Server, error) {
	opts := &server.Options{}

	// Initialize new server with options
	return server.NewServer(opts)
}

func lifecycle(lc fx.Lifecycle, ns *server.Server, ready ReadyChan) {
	lc.Append(fx.StartHook(func() {
		go func() {
			go ns.Start()

			// Wait for server to be ready for connections
			if !ns.ReadyForConnections(4 * time.Second) {
				panic("not ready for connection")
			}

			ready <- struct{}{}
		}()
	}))
}
