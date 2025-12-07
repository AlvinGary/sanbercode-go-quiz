package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// basic auth middleware
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == "admin" && password == "root" {
			c.Set("user", user)
			c.Next()
			return
		}
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}