package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func ProfileRoute(r *gin.RouterGroup, h *handler.ProfileHandler) {
	route := r.Group("/profile")
	{
		route.POST("", h.CreateProfile)
		route.GET("/:id", h.GetProfile)
		route.PUT("/:id", h.UpdateProfile)
		route.DELETE("/:id", h.DeleteProfile)
	}
}
