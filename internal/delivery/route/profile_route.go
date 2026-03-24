package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileRoute(r *gin.RouterGroup, h *handler.ProfileHandler, apiKey string) {
	route := r.Group("/profile")
	route.GET("/:id", h.GetProfile)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateProfile)
		protectedRoute.PUT("/:id", h.UpdateProfile)
		protectedRoute.DELETE("/:id", h.DeleteProfile)
	}
}
