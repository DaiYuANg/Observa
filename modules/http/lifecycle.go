package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func lifecycle(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.StartHook(func() {
		app.Listen(":8080")
	}))
}
