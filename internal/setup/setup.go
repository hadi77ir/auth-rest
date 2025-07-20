package setup

import (
	"auth-rest/internal/app"
	"auth-rest/internal/middlewares/globals"
	"auth-rest/internal/routers/auth"
	"auth-rest/internal/routers/users"
	"auth-rest/internal/swagger"
	"github.com/gofiber/fiber/v3"
)

type SetupFunc func(appGlobals *app.AppGlobals, router fiber.Router) error

func SetupHandlers(appGlobals *app.AppGlobals, app *fiber.App) error {
	funcs := []SetupFunc{
		// Add your Setup funcs here in order.

		// Middlewares
		// - REQUIRED: Globals Middleware
		globals.Setup,

		// v1 api setup
		apiV1Setup,

		// swagger setup
		swagger.Setup,
	}
	for _, setupFunc := range funcs {
		if err := setupFunc(appGlobals, app); err != nil {
			return err
		}
	}
	return nil
}

func apiV1Setup(appGlobals *app.AppGlobals, app fiber.Router) error {
	router := app.Group("/api/v1")

	funcs := []SetupFunc{
		// endpoints setup
		auth.Setup,
		users.Setup,
	}

	for _, setupFunc := range funcs {
		if err := setupFunc(appGlobals, router); err != nil {
			return err
		}
	}
	return nil
}
