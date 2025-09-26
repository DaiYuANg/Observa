package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/fx"
)

var Module = fx.Module("nats", fx.Provide(newClient))

func newClient() (*nats.Conn, jetstream.JetStream) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)

	return nc, js
}
