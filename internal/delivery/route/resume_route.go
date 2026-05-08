package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/go-chi/chi/v5"
)

func ResumeRoute(r chi.Router, h *handler.ResumeHandler) {
	r.Get("/resume/{profile_id}", h.GetResume)
}

