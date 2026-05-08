package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func EducationRoute(r fiber.Router, h *handler.EducationHandler, apiKey string) {
	route := r.Group("/education")
	route.Get("/:profile_id", h.GetEducation)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateEducation)
	protectedRoute.Put("/:education_id", h.UpdateEducation)
	protectedRoute.Delete("/:education_id", h.DeleteEducation)
}

