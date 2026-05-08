package middleware

import (
	"net/http"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(apiKey string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if c.Method() == http.MethodGet || c.Method() == http.MethodOptions {
			return c.Next()
		}

		clientKey := c.Get("X-API-Key")
		if clientKey != apiKey {
			return c.Status(http.StatusUnauthorized).JSON(dto.NoDataResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized: Invalid API Key",
			})
		}

		return c.Next()
	}
}

