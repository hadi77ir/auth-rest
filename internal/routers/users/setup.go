package users

import (
	"auth-rest/internal/app"
	"auth-rest/internal/middlewares/auth"
	"auth-rest/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"time"
)

func Setup(globals *app.AppGlobals, router fiber.Router) error {
	group := router.Group("/users")

	limit := limiter.New(limiter.Config{
		Storage:           storage.NewFromGlobals(globals, "limit-users"),
		LimiterMiddleware: limiter.SlidingWindow{},
		Max:               200,
		Expiration:        15 * time.Minute,
	})
	group.Use(limit)

	group.Use(auth.Middleware("admin"))
	group.Get("/", HandleList)
	group.Post("/", HandleCreate)
	group.Get("/:id", HandleSingleGet)
	group.Post("/:id", HandleSingleUpdate)
	group.Delete("/:id", HandleSingleDelete)

	router.Get("/user", HandleSingleGetSelf, auth.Middleware())
	router.Post("/user", HandleSingleUpdateSelf, auth.Middleware())

	return nil
}
