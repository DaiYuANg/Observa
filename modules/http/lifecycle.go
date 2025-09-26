package http

import (
	"github/DaiYuANg/Observa/internal/endpoint"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

func lifecycle(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.StartHook(func() {
		go app.Listen(":8080")
	}))
}

type RegisterControllerParam struct {
	fx.In
	Openapi     huma.API
	Controllers []endpoint.Endpoint `group:"endpoint"`
}

func registerEndpoint(param RegisterControllerParam) {
	endpoints, openapi := param.Controllers, param.Openapi
	lo.ForEach(endpoints, func(c endpoint.Endpoint, _ int) {
		c.RegisterRoute(openapi)
	})
}
