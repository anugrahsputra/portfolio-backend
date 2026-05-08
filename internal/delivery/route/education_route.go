package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func EducationRoute(r chi.Router, h *handler.EducationHandler, apiKey string) {
	r.Route("/education", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetEducation)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateEducation)
			r.Put("/{education_id}", h.UpdateEducation)
			r.Delete("/{education_id}", h.DeleteEducation)
		})
	})
}
