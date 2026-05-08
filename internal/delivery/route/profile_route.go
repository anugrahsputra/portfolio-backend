package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func ProfileRoute(r chi.Router, h *handler.ProfileHandler, apiKey string) {
	r.Route("/profile", func(r chi.Router) {
		r.Get("/{id}", h.GetProfile)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateProfile)
			r.Put("/{id}", h.UpdateProfile)
			r.Delete("/{id}", h.DeleteProfile)
		})
	})
}


