package dal

import (
	"auth-rest/internal/app"
	"errors"
	"github.com/gofiber/fiber/v3"
)

var ErrRecordNotFound = errors.New("record not found")

type Repositories interface {
	Users() UsersRepository
}

func FromContext(ctx fiber.Ctx) Repositories {
	return FromGlobals(app.FromContext(ctx))
}

const RepositoriesKey = "repo"

func FromGlobals(globals *app.AppGlobals) Repositories {
	return app.Value[Repositories](globals, RepositoriesKey)
}
