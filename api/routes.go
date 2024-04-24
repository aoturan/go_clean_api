package api

import "github.com/gofiber/fiber/v2"

func UserRouter(app fiber.Router, handler *UserHandler) {
	app.Get("/user", handler.HandleGetUser)
}

func SessionRouter(app fiber.Router, handler *SessionHandler) {
	app.Get("/session", handler.HandleGetSession)
}

func OrderRouter(app fiber.Router, handler *OrderHandler) {
	app.Get("/order", handler.HandleGetOrder)
}
