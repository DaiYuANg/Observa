package http

import (
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
	fx.Invoke(registerEndpoint, lifecycle),
)

func newServer() *fiber.App {
	app := fiber.New(fiber.Config{Prefork: false, EnablePrintRoutes: false, DisableStartupMessage: true})

	return app
}

func newOpenapi(app *fiber.App) huma.API {
	humaCfg := huma.DefaultConfig("observa", "1.0.0")
	humaCfg.DocsPath = "/"
	humaCfg.Servers = []*huma.Server{
		{URL: "http://127.0.0.1:" + "8080"},
		{URL: "http://localhost:" + "8080"},
	}
	return humafiber.New(app, humaCfg)
}
