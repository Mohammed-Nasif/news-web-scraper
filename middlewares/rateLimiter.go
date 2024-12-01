package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limiterGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  1,
	}

	store := memory.NewStore()

	instance := limiter.New(store, rate)

	limiterMiddleware := limiterGin.NewMiddleware(instance)

	return func(c *gin.Context) {
		limiterMiddleware(c)

		context, err := instance.Get(c.Request.Context(), c.ClientIP())
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
			return
		}

		if context.Reached {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests. Please try again later."})
			return
		}

		c.Next()
	}
}
