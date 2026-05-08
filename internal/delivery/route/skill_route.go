package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func SkillRoute(r chi.Router, h *handler.SkillHandler, apiKey string) {
	r.Route("/skill", func(r chi.Router) {
		r.Get("/{profile_id}", h.GetSkills)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(apiKey))
			r.Post("/", h.CreateSkill)
			r.Put("/{skill_id}", h.UpdateSkill)
			r.Delete("/{skill_id}", h.DeleteSkill)
		})
	})
}


