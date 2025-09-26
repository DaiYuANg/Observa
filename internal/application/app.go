package application

import "go.uber.org/fx"

type Application interface {
	GetApp() *fx.App
	GetOptions() fx.Option
}
