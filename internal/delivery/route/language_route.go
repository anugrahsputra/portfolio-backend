package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func LanguageRoute(r *gin.RouterGroup, h *handler.LanguageHandler) {
	route := r.Group("/language")
	{
		route.POST("", h.CreateLanguage)
		route.GET("/:profile_id", h.GetLanguages)
		route.PUT("/:language_id", h.UpdateLanguage)
		route.DELETE("/:language_id", h.DeleteLanguage)
	}
}
