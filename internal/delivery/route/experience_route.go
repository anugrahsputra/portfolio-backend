package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func ExperienceRoute(r fiber.Router, h *handler.ExperienceHandler, apiKey string) {
	route := r.Group("/experience")
	route.Get("/:profile_id", h.GetExperiences)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateExperience)
	protectedRoute.Put("/:experience_id", h.UpdateExperience)
	protectedRoute.Delete("/:experience_id", h.DeleteExperience)
}

