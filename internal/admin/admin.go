package admin

import (
	"github/DaiYuANg/Observa/internal/admin/router"
	"github/DaiYuANg/Observa/internal/injector"

	"go.uber.org/fx"
)

func NewAdmin() (*fx.App, error) {
	return injector.CreateContainer(
		router.Module,
	)
}
