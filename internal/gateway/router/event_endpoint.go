package router

import (
	"context"
	"github/DaiYuANg/Observa/internal/model"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nats-io/nats.go"
)

type EventEndpoint struct {
	nats *nats.Conn
}

func (n *EventEndpoint) RegisterRoute(openapi huma.API) {
	huma.Get[Test, model.Resp[any]](openapi, "/api/v1/event", n.OnEvent, huma.OperationTags("生成图片二维码"))
}

type Test struct {
}

func (n *EventEndpoint) OnEvent(ctx context.Context,
	parameter *Test) (*model.Resp[any], error) {

	return model.ReturnOk[model.R[any]](nil), nil
}
