package lib

import "github.com/gin-gonic/gin"

func ParamLimiter(cap int64, rate int64, key string) func(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	limiter := NewBucket(cap, rate)
	return func(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			if c.Query(key) != ""  {
				if !limiter.IsAccept() {
					c.AbortWithStatusJSON(429, gin.H{"message": "too many request"})
					return
				}
			}
			handlerFunc(c)
		}
	}
}
