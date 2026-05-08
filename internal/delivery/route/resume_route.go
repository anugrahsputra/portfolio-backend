package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gofiber/fiber/v3"
)

func ResumeRoute(r fiber.Router, h *handler.ResumeHandler) {
	route := r.Group("/resume")
	route.Get("/:profile_id", h.GetResume)
}
