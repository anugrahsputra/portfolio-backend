package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func ExperienceRoute(r chi.Router, h *handler.ExperienceHandler, apiKey string) {
	r.Route("/experience", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetExperiences)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateExperience)
			r.Put("/{experience_id}", h.UpdateExperience)
			r.Delete("/{experience_id}", h.DeleteExperience)
		})
	})
}


