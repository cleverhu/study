package middlewares

import (
	"github.com/gin-gonic/gin"
	"jtthinkStudy/casbinTest/lib"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.AbortWithStatusJSON(400, gin.H{"message": "require login"})
		}
		c.Set("user_name", token)
		c.Next()
	}
}

func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, _ := c.Get("user_name")
		ok, err := lib.E.Enforce(user, c.Request.RequestURI, c.Request.Method)
		if err != nil || !ok {
			c.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
		}
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			c.Abort()
		}
		c.Next()
	}
}

func Middlewares() []gin.HandlerFunc {
	ret := make([]gin.HandlerFunc, 0)
	ret = append(ret, Recovery(), CheckLogin(), RBAC())
	return ret
}
