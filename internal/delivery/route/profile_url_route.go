package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func ProfileUrlRoute(r *gin.RouterGroup, h *handler.ProfileUrlHandler) {
	route := r.Group("/profile-url")
	{
		route.POST("", h.CreateProfileUrl)
		route.GET("/:profile_url_id", h.GetProfileUrl)
		route.PUT("/:profile_url_id", h.UpdateProfileUrl)
		route.DELETE("/:profile_url_id", h.DeleteProfileUrl)
	}
}
