package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ExperienceRoute(r *gin.RouterGroup, h *handler.ExperienceHandler, apiKey string) {
	route := r.Group("/experience")
	route.GET("/:profile_id", h.GetExperiences)

	protected := route.Group("")
	protected.Use(middleware.AuthMiddleware(apiKey))
	{
		protected.POST("/", h.CreateExperience)
		protected.PUT("/:experience_id", h.UpdateExperience)
		protected.DELETE("/:experience_id", h.DeleteExperience)
	}
}
