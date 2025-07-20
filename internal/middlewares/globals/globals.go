package globals

import (
	"auth-rest/internal/app"
	"github.com/gofiber/fiber/v3"
)

func Middleware(globals *app.AppGlobals) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		ctx.Locals(app.AppGlobalsKey, globals)
		return ctx.Next()
	}
}

func Setup(globals *app.AppGlobals, router fiber.Router) error {
	router.Use(Middleware(globals))

	return nil
}
