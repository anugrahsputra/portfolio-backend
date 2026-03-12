package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func SkillRoute(r *gin.RouterGroup, h *handler.SkillHandler) {
	route := r.Group("/skill")
	{
		route.POST("", h.CreateSkill)
		route.GET("/:profile_id", h.GetSkills)
		route.PUT("/:skill_id", h.UpdateSkill)
		route.DELETE("/:skill_id", h.DeleteSkill)
	}
}
