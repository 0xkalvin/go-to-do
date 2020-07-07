package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/lucsky/cuid.v1"
)

func AddRequestId()  gin.HandlerFunc {
	return func(c *gin.Context){
		id := cuid.New()
		c.Writer.Header().Set("X-Request-Id", id)
		c.Next()
	}
}