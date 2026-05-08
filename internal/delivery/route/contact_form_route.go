package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func ContactFormRoute(r chi.Router, h *handler.ContactFormHandler, apiKey string) {
	r.With(middleware.AuthMiddleware(apiKey)).Post("/send-email", h.SendMail)
}


