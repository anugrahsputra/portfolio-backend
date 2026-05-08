package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func SkillRoute(r fiber.Router, h *handler.SkillHandler, apiKey string) {
	route := r.Group("/skill")
	route.Get("/:profile_id", h.GetSkills)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateSkill)
	protectedRoute.Put("/:skill_id", h.UpdateSkill)
	protectedRoute.Delete("/:skill_id", h.DeleteSkill)
}

