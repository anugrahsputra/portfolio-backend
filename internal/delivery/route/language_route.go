package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func LanguageRoute(r fiber.Router, h *handler.LanguageHandler, apiKey string) {
	route := r.Group("/language")
	route.Get("/:profile_id", h.GetLanguages)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateLanguage)
	protectedRoute.Put("/:language_id", h.UpdateLanguage)
	protectedRoute.Delete("/:language_id", h.DeleteLanguage)
}

