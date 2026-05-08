package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func ProjectRoute(r chi.Router, h *handler.ProjectHandler, apiKey string) {
	r.Route("/project", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetProjects)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateProject)
			r.Put("/{project_id}", h.UpdateProject)
			r.Delete("/{project_id}", h.DeleteProject)
		})
	})
}


