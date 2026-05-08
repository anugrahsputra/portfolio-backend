package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func ProjectRoute(r fiber.Router, h *handler.ProjectHandler, apiKey string) {
	route := r.Group("/project")
	route.Get("/:profile_id", h.GetProjects)

	protectedRoute := route.Group("/", middleware.AuthMiddleware(apiKey))
	protectedRoute.Post("/", h.CreateProject)
	protectedRoute.Put("/:project_id", h.UpdateProject)
	protectedRoute.Delete("/:project_id", h.DeleteProject)
}

