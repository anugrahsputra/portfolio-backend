package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ExperienceRoute(r *gin.RouterGroup, h *handler.ExperienceHandler, apiKey string) {
	route := r.Group("/experience")
	route.GET("/:profile_id", h.GetExperiences)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateExperience)
		protectedRoute.PUT("/:experience_id", h.UpdateExperience)
		protectedRoute.DELETE("/:experience_id", h.DeleteExperience)
	}
}
