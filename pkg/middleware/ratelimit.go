package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key := "ratelimit:sendemail:" + c.ClientIP()

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, dto.NoDataResponse{
				Status:  http.StatusTooManyRequests,
				Message: "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
