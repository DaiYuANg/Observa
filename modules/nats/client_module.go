package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/fx"
)

var ClientModule = fx.Module("nats", fx.Provide(newClient, newJetStream))

func newClient() (*nats.Conn, error) {
	return nats.Connect(nats.DefaultURL)
}

func newJetStream(nc *nats.Conn) (jetstream.JetStream, error) {
	return jetstream.New(nc)
}
