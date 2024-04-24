package api

import (
	"github.com/aoturan/go_clean_api/pkg/session"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type SessionHandler struct {
	sessionService session.Service
}

func NewSessionHandler(sessionService session.Service) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) HandleGetSession(c *fiber.Ctx) error {
	log.Info().Msg("user started")
	return c.JSON("hello world")
}
