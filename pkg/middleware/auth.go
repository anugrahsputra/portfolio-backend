package middleware

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		clientKey := c.GetHeader("X-API-Key")
		if clientKey != apiKey {
			c.JSON(http.StatusUnauthorized, dto.NoDataResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized: Invalid API Key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
