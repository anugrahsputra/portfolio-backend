package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileUrlRoute(r *gin.RouterGroup, h *handler.ProfileUrlHandler, apiKey string) {
	route := r.Group("/profile-url")
	route.GET("/:profile_id", h.GetProfileURL)
	route.GET("/url/:profile_url_id", h.GetProfileUrlByID)

	protected := route.Group("")
	protected.Use(middleware.AuthMiddleware(apiKey))
	{
		protected.POST("/", h.CreateProfileUrl)
		protected.PUT("/:profile_url_id", h.UpdateProfileUrl)
		protected.DELETE("/:profile_url_id", h.DeleteProfileUrl)
	}
}
