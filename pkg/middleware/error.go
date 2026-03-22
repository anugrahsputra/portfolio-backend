package middleware

import (
	"fmt"
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Get the last error
			err := c.Errors.Last()

			// Check if it's a validation error or other
			status := http.StatusInternalServerError
			message := "Internal Server Error"

			// We can refine this to handle different error types if needed
			// But for now, we just return a generic message to the client
			// while logging the actual error.

			fmt.Printf("API Error: %v\n", err.Err)

			c.JSON(status, dto.NoDataResponse{
				Status:  status,
				Message: message,
			})
			c.Abort()
		}
	}
}

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
