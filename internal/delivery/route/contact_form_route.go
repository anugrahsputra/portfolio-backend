package route

import (
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ContactFormRoute(r *gin.RouterGroup, h *handler.ContactFormHandler, apiKey string) {
	route := r.Group("/send-email")
	route.Use(middleware.AuthMiddleware(apiKey))
	{
		route.POST("/", h.SendMail)
	}
}
