package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func newRateLimitRouter(rdb *redis.Client, limit int, window time.Duration) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/send-email", middleware.RateLimitMiddleware(rdb, limit, window), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"sent": true})
	})
	return r
}

func doPost(r *gin.Engine, ip string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/send-email", nil)
	req.RemoteAddr = ip + ":12345"
	r.ServeHTTP(w, req)
	return w
}

func TestRateLimitMiddleware_AllowsUpToLimitThenBlocks(t *testing.T) {
	rdb, _ := newTestRedis(t)
	r := newRateLimitRouter(rdb, 3, time.Minute)

	for i := 0; i < 3; i++ {
		w := doPost(r, "1.2.3.4")
		assert.Equal(t, http.StatusOK, w.Code, "request %d should be allowed", i+1)
	}

	w := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusTooManyRequests, w.Code, "4th request in window should be blocked")
}

func TestRateLimitMiddleware_DifferentIPsTrackedSeparately(t *testing.T) {
	rdb, _ := newTestRedis(t)
	r := newRateLimitRouter(rdb, 1, time.Minute)

	w1 := doPost(r, "1.1.1.1")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := doPost(r, "2.2.2.2")
	assert.Equal(t, http.StatusOK, w2.Code, "different IP must have its own counter")
}

func TestRateLimitMiddleware_ResetsAfterWindowExpires(t *testing.T) {
	rdb, mr := newTestRedis(t)
	r := newRateLimitRouter(rdb, 1, time.Minute)

	w1 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusTooManyRequests, w2.Code)

	mr.FastForward(time.Minute + time.Second)

	w3 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w3.Code, "counter must reset once the window elapses")
}

func TestRateLimitMiddleware_FailsOpenWhenRedisDown(t *testing.T) {
	rdb, mr := newTestRedis(t)
	mr.Close()

	r := newRateLimitRouter(rdb, 1, time.Minute)

	w := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w.Code, "request must succeed when Redis is unreachable")
}
