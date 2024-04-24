package api

import (
	"github.com/aoturan/go_clean_api/pkg/order"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type OrderHandler struct {
	orderService order.Service
}

func NewOrderHandler(orderService order.Service) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) HandleGetOrder(c *fiber.Ctx) error {
	log.Info().Msg("order started")
	return c.JSON("hello world")
}
