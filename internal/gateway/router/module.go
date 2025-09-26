package router

import (
	"github/DaiYuANg/Observa/internal/endpoint"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

var Module = fx.Module("router",
	fx.Provide(
		endpoint.Annotation(newEventEndPoint),
	),
)

func newEventEndPoint(c *nats.Conn) *EventEndpoint {
	return &EventEndpoint{
		nats: c,
	}
}
