package auth

import (
	"auth-rest/internal/app"
	"auth-rest/internal/middlewares/auth"
	"auth-rest/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"time"
)

func Setup(globals *app.AppGlobals, router fiber.Router) error {
	group := router.Group("/auth")

	limit := limiter.New(limiter.Config{
		Storage:           storage.NewFromGlobals(globals, "limit-auth"),
		LimiterMiddleware: limiter.SlidingWindow{},
		Max:               20,
		Expiration:        15 * time.Minute,
	})
	group.Use(limit)

	group.Post("/request-otp", HandleRequestOTP)
	group.Post("/login", HandleLogin)
	group.Post("/logout", HandleLogout, auth.Middleware())

	return nil
}
