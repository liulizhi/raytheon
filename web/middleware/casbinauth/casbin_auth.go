package casbinauth

import (
	"net/http"
	"raytheon/pkg/e"
	"raytheon/web/middleware/jwtauth"

	"github.com/gin-gonic/gin"
)

// AuthCheckRole role check
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*jwtauth.CustomClaims)
		role := claims.Role
		res, err := e.CasbinEnforcer.EnforceSafe(role, claims.Tenant, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "错误消息" + err.Error(),
			})
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "很抱歉您没有此权限",
			})
			c.Abort()
			return
		}
	}
}
