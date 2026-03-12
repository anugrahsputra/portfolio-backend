package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func EducationRoute(r *gin.RouterGroup, h *handler.EducationHandler) {
	route := r.Group("/education")
	{
		route.POST("", h.CreateEducation)
		route.GET("/:profile_id", h.GetEducation)
		route.PUT("/:education_id", h.UpdateEducation)
		route.DELETE("/:education_id", h.DeleteEducation)
	}
}
