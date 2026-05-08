package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func ProfileUrlRoute(r chi.Router, h *handler.ProfileUrlHandler, apiKey string) {
	r.Route("/profile-url", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetProfileURL)
		r.Get("/url/{profile_url_id}", h.GetProfileUrlByID)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateProfileUrl)
			r.Put("/{profile_url_id}", h.UpdateProfileUrl)
			r.Delete("/{profile_url_id}", h.DeleteProfileUrl)
		})
	})
}


