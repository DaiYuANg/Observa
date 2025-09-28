package router

import (
	"context"
	"encoding/json"
	"github/DaiYuANg/Observa/internal/model"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nats-io/nats.go"
)

type EventEndpoint struct {
	nats *nats.Conn
}

func (n *EventEndpoint) RegisterRoute(openapi huma.API) {
	huma.Post[EventRequest, model.Resp[any]](
		openapi,
		"/api/v1/event",
		n.OnEvent,
		huma.OperationTags("Receive SDK event"))
}

// EventRequest SDK 上报的事件结构
type EventRequest struct {
	Body struct {
		AppID      string                 `json:"app_id"`      // SDK 所属应用
		EventType  string                 `json:"event_type"`  // 事件类型，例如 "error"
		Timestamp  int64                  `json:"timestamp"`   // Unix 时间戳，毫秒
		Payload    map[string]interface{} `json:"payload"`     // 自定义事件数据
		DeviceInfo map[string]string      `json:"device_info"` // 设备信息，可选
		UserID     string                 `json:"user_id"`     // 用户 ID，可选
	}
}

func (n *EventEndpoint) OnEvent(
	ctx context.Context,
	parameter *EventRequest,
) (*model.Resp[any], error) {
	// 如果没有 timestamp，填入当前时间
	var body = parameter.Body
	if body.Timestamp == 0 {
		body.Timestamp = time.Now().UnixMilli()
	}

	// 序列化事件为 JSON
	data, err := json.Marshal(parameter)
	if err != nil {
		return nil, err
	}

	// 发布到 NATS 主题，例如 "observa.events"
	if err := n.nats.Publish("observa.events", data); err != nil {
		return nil, err
	}
	return model.ReturnOk[model.R[any]](nil), nil
}
