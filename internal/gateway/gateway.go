package gateway

import (
	"github/DaiYuANg/Observa/internal/gateway/router"
	"github/DaiYuANg/Observa/internal/injector"
	"github/DaiYuANg/Observa/modules/http"
	"github/DaiYuANg/Observa/modules/nats"

	"go.uber.org/fx"
)

type Application struct {
	options []fx.Option
	app     *fx.App
}

func (a Application) GetApp() *fx.App {
	return a.app
}

func (a Application) GetOptions() []fx.Option {
	return a.options
}

func NewGateway() (*Application, error) {
	app, options, err := injector.CreateContainer(router.Module, nats.ClientModule, http.Module)
	if err != nil {
		return nil, err
	}
	return &Application{
		app:     app,
		options: options,
	}, nil
}
