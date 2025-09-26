package router

import (
	"github/DaiYuANg/Observa/internal/endpoint"

	"go.uber.org/fx"
)

var Module = fx.Module("router",
	fx.Provide(
		endpoint.Annotation(newEventEndPoint),
	),
)

func newEventEndPoint() *EventEndpoint {
	return &EventEndpoint{}
}
