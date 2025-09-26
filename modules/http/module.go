package http

import (
	"github/DaiYuANg/Observa/modules/config"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(
		newServer,
		newOpenapi,
	),
	fx.Invoke(lifecycle),
)

func newServer() *fiber.App {
	app := fiber.New()

	return app
}

func newOpenapi(app *fiber.App, cfg *config.Config) huma.API {
	humaCfg := huma.DefaultConfig("captcha service", "1.0.0")
	humaCfg.DocsPath = "/"
	humaCfg.Servers = []*huma.Server{
		{URL: "http://127.0.0.1:" + cfg.Http.GetPort()},
		{URL: "http://localhost:" + cfg.Http.GetPort()},
	}
	return humafiber.New(app, humaCfg)
}
