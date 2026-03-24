package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func EducationRoute(r *gin.RouterGroup, h *handler.EducationHandler, apiKey string) {
	route := r.Group("/education")
	route.GET("/:profile_id", h.GetEducation)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateEducation)
		protectedRoute.PUT("/:education_id", h.UpdateEducation)
		protectedRoute.DELETE("/:education_id", h.DeleteEducation)
	}
}
