package auth

import (
	"auth-rest/internal/middlewares/auth"
	"auth-rest/internal/modules/jwt"
	"auth-rest/internal/routers/common"
	"github.com/gofiber/fiber/v3"
)

// HandleLogout is the logout handler
//
// @Summary      Logout
// @Description  Revokes current access token
// @Tags         auth
// @Produce      json
// @Security     BearerToken
// @Success      200  {object}  common.SuccessResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /auth/logout [post]
func HandleLogout(ctx fiber.Ctx) error {
	manager := jwt.FromContext(ctx)
	claims := auth.ClaimsFromContext(ctx)
	err := manager.Revoke(claims.TokenID)
	if err != nil {
		return err
	}
	return ctx.JSON(&common.SuccessResponse{Success: true})
}
