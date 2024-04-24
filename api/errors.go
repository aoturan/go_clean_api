package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiError Error
	if errors.As(err, &apiError) {
		return c.Status(apiError.Code).JSON(apiError)
	}

	apiError = NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

// Error implements Error interface
func (e Error) Error() string {
	return e.Err
}
