package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type cachedResponse struct {
	Status      int    `json:"status"`
	ContentType string `json:"content_type"`
	Body        []byte `json:"body"`
}

type bodyCapturingWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyCapturingWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CacheMiddleware(rdb *redis.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		ctx := context.Background()
		key := "cache:" + c.Request.URL.RequestURI()

		cached, err := rdb.Get(ctx, key).Bytes()
		if err == nil {
			var resp cachedResponse
			if json.Unmarshal(cached, &resp) == nil {
				c.Header("Content-Type", resp.ContentType)
				c.Data(resp.Status, resp.ContentType, resp.Body)
				c.Abort()
				return
			}
		}

		writer := &bodyCapturingWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = writer

		c.Next()

		if writer.Status() != http.StatusOK {
			return
		}

		encoded, err := json.Marshal(cachedResponse{
			Status:      writer.Status(),
			ContentType: writer.Header().Get("Content-Type"),
			Body:        writer.body.Bytes(),
		})
		if err != nil {
			return
		}

		rdb.Set(ctx, key, encoded, ttl)
	}
}
