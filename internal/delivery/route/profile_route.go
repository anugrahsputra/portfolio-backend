package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func ProfileRoute(r fiber.Router, h *handler.ProfileHandler, apiKey string) {
	route := r.Group("/profile")
	route.Get("/:id", h.GetProfile)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateProfile)
	protectedRoute.Put("/:id", h.UpdateProfile)
	protectedRoute.Delete("/:id", h.DeleteProfile)
}

