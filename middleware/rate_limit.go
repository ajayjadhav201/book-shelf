package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(r rate.Limit, b int) gin.HandlerFunc {
	fmt.Println("ajaj rate limiter called")
	limiter := rate.NewLimiter(r, b)
	fmt.Println("ajaj rate limiter created")
	return func(c *gin.Context) {
		if !limiter.AllowN(time.Now(), 1) {
			c.String(http.StatusTooManyRequests, "Rate limit exceeded")
			c.Abort()
			return
		}
		c.Next()
	}
}
