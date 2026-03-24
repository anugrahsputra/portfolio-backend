package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func LanguageRoute(r *gin.RouterGroup, h *handler.LanguageHandler, apiKey string) {
	route := r.Group("/language")
	route.GET("/:profile_id", h.GetLanguages)

	protectedRoute := route.Group("")
	protectedRoute.Use(middleware.AuthMiddleware(apiKey))
	{
		protectedRoute.POST("", h.CreateLanguage)
		protectedRoute.PUT("/:language_id", h.UpdateLanguage)
		protectedRoute.DELETE("/:language_id", h.DeleteLanguage)
	}
}
