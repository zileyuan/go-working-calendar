package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CrossAccessAllow 可以跨域访问的中间件
func CrossAccessAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")                                                                            //允许访问所有域
		header.Add("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")                                             //允许访问的方法
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With") //header的类型
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
