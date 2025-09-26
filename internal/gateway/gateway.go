package gateway

import (
	"github/DaiYuANg/Observa/internal/gateway/router"
	"github/DaiYuANg/Observa/internal/injector"
	"github/DaiYuANg/Observa/modules/http"
	"github/DaiYuANg/Observa/modules/nats"

	"go.uber.org/fx"
)

func NewGateway() (*fx.App, error) {
	return injector.CreateContainer(router.Module, nats.Module, http.Module)
}
