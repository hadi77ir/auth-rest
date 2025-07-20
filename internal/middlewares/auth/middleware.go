package auth

import (
	"auth-rest/internal/dal"
	"auth-rest/internal/modules/auth"
	"auth-rest/internal/modules/jwt"
	"auth-rest/internal/modules/users"
	"github.com/gofiber/fiber/v3"
	"strconv"
	"strings"
)

const ClaimsKey = "jwt-claims"
const UserKey = "auth-user"

func Middleware(roles ...string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		manager := jwt.FromContext(ctx)
		token := TokenFromHeader(ctx)
		if token == "" {
			return fiber.ErrUnauthorized
		}
		claims, err := manager.Validate(token)
		if err != nil {
			return err
		}
		// get user from db
		uid, err := strconv.Atoi(claims.Subject)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		repo := dal.FromContext(ctx)
		user, err := users.UserByID(repo.Users(), uid)
		if !auth.HasPermission(user.Role, roles) {
			return fiber.ErrForbidden
		}
		ctx.Locals(UserKey, user)
		ctx.Locals(ClaimsKey, claims)
		return ctx.Next()
	}
}

func TokenFromHeader(ctx fiber.Ctx) string {
	tokenHeader := ctx.Get(fiber.HeaderAuthorization)
	if tokenHeader == "" {
		return ""
	}
	if !strings.HasPrefix(tokenHeader, "Bearer ") {
		return ""
	}
	token := tokenHeader[len("Bearer "):]
	return token
}

func UserFromContext(ctx fiber.Ctx) users.UserModel {
	user := ctx.Locals(UserKey)
	if user == nil {
		return users.UserModel{}
	}
	return user.(users.UserModel)
}

func ClaimsFromContext(ctx fiber.Ctx) *jwt.TokenClaims {
	claims := ctx.Locals(ClaimsKey)
	if claims == nil {
		return nil
	}
	return claims.(*jwt.TokenClaims)
}
