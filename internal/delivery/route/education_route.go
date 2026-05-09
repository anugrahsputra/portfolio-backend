package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func EducationRoute(r *gin.RouterGroup, h *handler.EducationHandler, apiKey string) {
	route := r.Group("/education")
	route.GET("/:profile_id", h.GetEducation)

	protected := route.Group("")
	protected.Use(middleware.AuthMiddleware(apiKey))
	{
		protected.POST("/", h.CreateEducation)
		protected.PUT("/:education_id", h.UpdateEducation)
		protected.DELETE("/:education_id", h.DeleteEducation)
	}
}
