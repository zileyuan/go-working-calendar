package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zileyuan/go-working-calendar/router"
	"github.com/zileyuan/go-working-calendar/util"
)

// AuthToken 验证Token的中间件
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		pass := false
		fullPath := c.FullPath()
		util.Infof("Route full path: %v\n", fullPath)
		auth := router.AuthTokenRoutes[fullPath]
		if auth {
			var tokenString = c.Request.Header.Get("authorization")
			if tokenString != "" {
				valid, data := util.VerifyToken(tokenString)
				if valid {
					for k, v := range data {
						c.Request.Header.Set(k, v.(string))
					}
					pass = true
				}
			}
		} else {
			pass = true
		}
		if pass {
			c.Next()
		} else {
			uri := c.Request.RequestURI
			util.Infof("URI(%v) is blocked by AuthToken from FullPath(%v)\n", uri, fullPath)
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
