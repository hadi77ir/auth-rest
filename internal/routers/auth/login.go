package auth

import (
	"auth-rest/internal/dal"
	"auth-rest/internal/modules/auth"
	"auth-rest/internal/modules/jwt"
	"auth-rest/internal/modules/sms"
	"auth-rest/internal/routers/common"
	"auth-rest/internal/utils"
	goerrors "errors"
	"github.com/gofiber/fiber/v3"
	"strconv"
	"time"
)

// HandleRequestOTP is handler for requesting OTP SMS
//
// @Summary      Request OTP
// @Description  Sends an SMS containing OTP
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginInitRequest true "Login init request"
// @Success      200  {object}  common.SuccessResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /auth/request-otp [post]
func HandleRequestOTP(ctx fiber.Ctx) error {
	repos := dal.FromContext(ctx)
	var request LoginRequest
	err := utils.DecodeJSONBody(ctx, &request)
	if err != nil {
		return err
	}
	if !utils.IsValidPhone(request.Phone) {
		return fiber.ErrBadRequest
	}
	smsProvider := sms.FromContext(ctx)
	_, err = auth.InitiateUserLogin(repos.Users(), smsProvider, request.Phone)
	if err != nil {
		return err
	}
	return ctx.JSON(&common.SuccessResponse{Success: true})
}

// HandleLogin is the login handler
//
// @Summary      Login by OTP
// @Description  Returns an access token upon success
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login request"
// @Success      200  {object}  LoginResponse
// @Failure      401  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /auth/login [post]
func HandleLogin(ctx fiber.Ctx) error {
	repos := dal.FromContext(ctx)
	var request LoginRequest
	err := utils.DecodeJSONBody(ctx, &request)
	if err != nil {
		return err
	}
	if !utils.IsValidPhone(request.Phone) {
		return fiber.ErrBadRequest
	}
	user, err := auth.VerifyOTP(repos.Users(), request.Phone, request.OTP)
	if err != nil {
		if goerrors.Is(err, auth.ErrInvalidOTP) {
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			}
		}
		return err
	}
	manager := jwt.FromContext(ctx)
	token, err := manager.Generate(strconv.Itoa(int(user.ID)), time.Now().Add(12*time.Hour))
	if err != nil {
		return err
	}
	return ctx.JSON(&LoginResponse{Token: token})
}
