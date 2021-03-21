package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {

		s := c.GetHeader("userName")
		d := c.GetHeader("domain")
		p := c.Request.URL.Path
		m := c.Request.Method

		if b, err := e.Enforce(s, d, p, m); err != nil {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
			return
		} else if !b {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
			return
		}
		fmt.Println("恭喜您,权限验证通过")
		c.Next()
	}
}
