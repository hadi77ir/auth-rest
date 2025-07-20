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

	limitMain := limiter.New(limiter.Config{
		Storage:           storage.NewFromGlobals(globals, "limit-auth"),
		LimiterMiddleware: limiter.SlidingWindow{},
		Max:               20,
		Expiration:        15 * time.Minute,
	})
	group.Use(limitMain)
	limitOTP := limiter.New(limiter.Config{
		Storage:           storage.NewFromGlobals(globals, "limit-otp"),
		LimiterMiddleware: limiter.FixedWindow{},
		Max:               1,
		Expiration:        2 * time.Minute,
	})

	group.Post("/request-otp", HandleRequestOTP, limitOTP)
	group.Post("/login", HandleLogin, limitMain)
	group.Post("/logout", HandleLogout, limitMain, auth.Middleware())

	return nil
}
