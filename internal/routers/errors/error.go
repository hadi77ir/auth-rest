package errors

import (
	"auth-rest/internal/dal"
	"auth-rest/internal/routers/common"
	"errors"
	"github.com/gofiber/fiber/v3"
)

func HandleError(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	var fiberError *fiber.Error
	if errors.As(err, &fiberError) && fiberError.Code != 0 {
		c.Status(fiberError.Code)
	} else if errors.Is(err, dal.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
	} else {
		c.Status(fiber.StatusBadRequest)
	}
	return c.JSON(&common.ErrorResponse{
		Error: err.Error(),
	})
}
