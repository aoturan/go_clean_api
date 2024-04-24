package api

import (
	"github.com/aoturan/go_clean_api/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	log.Info().Msg("user started")
	return c.JSON("hello world")
}
