package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileUrlRoute(r *gin.RouterGroup, h *handler.ProfileUrlHandler, apiKey string) {
	route := r.Group("/profile-url")
	route.GET("/:profile_url_id", h.GetProfileUrl)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateProfileUrl)
		protectedRoute.PUT("/:profile_url_id", h.UpdateProfileUrl)
		protectedRoute.DELETE("/:profile_url_id", h.DeleteProfileUrl)
	}
}
