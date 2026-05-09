package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func LanguageRoute(r *gin.RouterGroup, h *handler.LanguageHandler, apiKey string) {
	route := r.Group("/language")
	route.GET("/:profile_id", h.GetLanguages)

	protected := route.Group("")
	protected.Use(middleware.AuthMiddleware(apiKey))
	{
		protected.POST("/", h.CreateLanguage)
		protected.PUT("/:language_id", h.UpdateLanguage)
		protected.DELETE("/:language_id", h.DeleteLanguage)
	}
}
