package jwt

import (
	"auth-rest/internal/app"
	"auth-rest/internal/config"
	"auth-rest/internal/storage"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type TokenClaims struct {
	TokenID string `json:"token_id"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	cfg        *config.JWTConfig
	signMethod jwt.SigningMethod
	storage    fiber.Storage
}

func (m *TokenManager) Generate(subject string, expiration time.Time) (string, error) {
	tokenID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(m.signMethod, &TokenClaims{
		TokenID: tokenID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Issuer:    m.cfg.Issuer,
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	tokenString, err := token.SignedString(m.cfg.Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *TokenManager) Validate(tokenStr string) (*TokenClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != m.signMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return m.cfg.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("token has invalid claims")
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}
	if claims.Issuer != m.cfg.Issuer {
		return nil, errors.New("token has invalid issuer")
	}
	if m.storage != nil {
		val, err := m.storage.Get(getRevokedTokenKey(claims.TokenID))
		if err != nil {
			return nil, err
		}
		if len(val) == 1 && val[0] == 1 {
			return nil, errors.New("token is revoked")
		}
	}
	return claims, nil
}
func (m *TokenManager) Revoke(tokenId string) error {
	if m.storage != nil {
		_ = m.storage.Set(getRevokedTokenKey(tokenId), []byte{1}, 0)
	}
	return nil
}

func getRevokedTokenKey(tokenId string) string {
	return "revoked-token-" + tokenId
}

func NewTokenManager(cfg *config.JWTConfig, storage fiber.Storage) (*TokenManager, error) {
	if cfg == nil {
		return nil, errors.New("config is nil")
	}
	return &TokenManager{
		cfg:        cfg,
		signMethod: jwt.GetSigningMethod(cfg.Algorithm),
		storage:    storage,
	}, nil
}

const ManagerKey = "jwt-manager"

func Setup(globals *app.AppGlobals) error {
	cfg := config.FromGlobals(globals)
	store := storage.NewFromGlobals(globals, "revoked-jwts")
	manager, err := NewTokenManager(cfg.JWT, store)
	if err != nil {
		return err
	}
	globals.Set(ManagerKey, manager)
	return nil
}

func FromContext(ctx fiber.Ctx) *TokenManager {
	return app.Value[*TokenManager](app.FromContext(ctx), ManagerKey)
}
