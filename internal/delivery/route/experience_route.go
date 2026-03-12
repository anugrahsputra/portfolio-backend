package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func ExperienceRoute(r *gin.RouterGroup, h *handler.ExperienceHandler) {
	route := r.Group("/experience")
	{
		route.POST("", h.CreateExperience)
		route.GET("/:profile_id", h.GetExperiences)
		route.PUT("/:experience_id", h.UpdateExperience)
		route.DELETE("/:experience_id", h.DeleteExperience)
	}
}
