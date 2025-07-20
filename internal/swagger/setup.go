package swagger

import (
	"auth-rest/internal/app"
	_ "auth-rest/internal/docs"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger"
)

//	@title			Auth-Rest
//	@description	Simple Auth Server
//	@version		1.0

//	@host		localhost:4000
//	@BasePath	/api/v1

//	@securityDefinitions.bearer	BearerToken
//	@in							header
//	@name						Authorization

func Setup(appGlobals *app.AppGlobals, router fiber.Router) error {
	router.Get("/swagger/*", swagger.HandlerDefault)
	return nil
}
