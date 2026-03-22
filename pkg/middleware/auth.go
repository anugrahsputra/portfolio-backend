package middleware

import (
	"net/http"
	"os"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only check auth for non-GET requests (e.g., POST, PUT, DELETE)
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			// In case the API_KEY is not set, we can either allow it in development
			// or fail safe by blocking everything in production.
			env := os.Getenv("ENV")
			if env != "development" {
				c.JSON(http.StatusUnauthorized, dto.NoDataResponse{
					Status:  http.StatusUnauthorized,
					Message: "API Key not configured on server",
				})
				c.Abort()
				return
			}
		}

		clientKey := c.GetHeader("X-API-KEY")
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
