package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func newTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	t.Cleanup(mr.Close)

	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	return client, mr
}

func newTestRouter(rdb *redis.Client, hitCount *int) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.CacheMiddleware(rdb, time.Minute))
	r.GET("/thing", func(c *gin.Context) {
		*hitCount++
		c.JSON(http.StatusOK, gin.H{"count": *hitCount})
	})
	r.POST("/thing", func(c *gin.Context) {
		*hitCount++
		c.JSON(http.StatusOK, gin.H{"count": *hitCount})
	})
	return r
}

func TestCacheMiddleware_SecondGetIsServedFromCache(t *testing.T) {
	rdb, _ := newTestRedis(t)
	hits := 0
	r := newTestRouter(rdb, &hits)

	w1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.JSONEq(t, `{"count":1}`, w1.Body.String())

	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.JSONEq(t, `{"count":1}`, w2.Body.String())
	assert.Equal(t, 1, hits, "handler should only run once; second GET served from cache")
}

func TestCacheMiddleware_NonGetBypassesCache(t *testing.T) {
	rdb, _ := newTestRedis(t)
	hits := 0
	r := newTestRouter(rdb, &hits)

	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/thing", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	assert.Equal(t, 2, hits, "POST requests must never be cached")
}

func TestCacheMiddleware_FailsOpenWhenRedisDown(t *testing.T) {
	rdb, mr := newTestRedis(t)
	mr.Close()

	hits := 0
	r := newTestRouter(rdb, &hits)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, hits, "request must still succeed via the handler when Redis is unreachable")
}
