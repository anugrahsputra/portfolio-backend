package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func ProfileUrlRoute(r fiber.Router, h *handler.ProfileUrlHandler, apiKey string) {
	route := r.Group("/profile-url")
	route.Get("/:profile_id", h.GetProfileURL)
	route.Get("/url/:profile_url_id", h.GetProfileUrlByID)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateProfileUrl)
	protectedRoute.Put("/:profile_url_id", h.UpdateProfileUrl)
	protectedRoute.Delete("/:profile_url_id", h.DeleteProfileUrl)
}

