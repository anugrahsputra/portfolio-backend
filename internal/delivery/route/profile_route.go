package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileRoute(r *gin.RouterGroup, h *handler.ProfileHandler, apiKey string) {
	route := r.Group("/profile")
	route.GET("/:id", h.GetProfile)

	protected := route.Group("")
	protected.Use(middleware.AuthMiddleware(apiKey))
	{
		protected.POST("/", h.CreateProfile)
		protected.PUT("/:id", h.UpdateProfile)
		protected.DELETE("/:id", h.DeleteProfile)
	}
}
