package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoute(r *gin.RouterGroup, h *handler.ProjectHandler, apiKey string) {
	route := r.Group("/project")
	route.GET("/:profile_id", h.GetProjects)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateProject)
		protectedRoute.PUT("/:project_id", h.UpdateProject)
		protectedRoute.DELETE("/:project_id", h.DeleteProject)
	}
}



