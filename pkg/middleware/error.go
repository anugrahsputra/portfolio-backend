package middleware

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic recovered: %v\n", err)
				c.JSON(http.StatusInternalServerError, dto.NoDataResponse{
					Status:  http.StatusInternalServerError,
					Message: "A serious error occurred on the server.",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
