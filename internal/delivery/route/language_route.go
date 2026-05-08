package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func LanguageRoute(r chi.Router, h *handler.LanguageHandler, apiKey string) {
	r.Route("/language", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetLanguages)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateLanguage)
			r.Put("/{language_id}", h.UpdateLanguage)
			r.Delete("/{language_id}", h.DeleteLanguage)
		})
	})
}


