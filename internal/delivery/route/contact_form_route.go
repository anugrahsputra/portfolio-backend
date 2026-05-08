package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func ContactFormRoute(r fiber.Router, h *handler.ContactFormHandler, apiKey string) {
	route := r.Group("/send-email", middleware.AuthMiddleware(apiKey))
	route.Post("/", h.SendMail)
}

