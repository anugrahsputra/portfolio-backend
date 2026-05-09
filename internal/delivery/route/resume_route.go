package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func ResumeRoute(r *gin.RouterGroup, h *handler.ResumeHandler) {
	route := r.Group("/resume")
	route.GET("/:profile_id", h.GetResume)
}


