package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SkillRoute(r *gin.RouterGroup, h *handler.SkillHandler, apiKey string) {
	route := r.Group("/skill")
	route.GET("/:profile_id", h.GetSkills)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateSkill)
		protectedRoute.PUT("/:skill_id", h.UpdateSkill)
		protectedRoute.DELETE("/:skill_id", h.DeleteSkill)
	}
}
