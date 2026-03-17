package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func ProjectRoute(r *gin.RouterGroup, h *handler.ProjectHandler) {
	route := r.Group("/project")
	{
		route.POST("", h.CreateProject)
		route.GET("/:profile_id", h.GetProjects)
		route.PUT("/:project_id", h.UpdateProject)
		route.DELETE("/:project_id", h.DeleteProject)
	}
}
