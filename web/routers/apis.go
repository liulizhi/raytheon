package routers

import (
	"raytheon/web/controllers/v1/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterRouter register all router
func RegisterRouter(router *gin.Engine) {
	// Login
	router.Use(cors.Default())

	apiVersion := router.Group("api/v1")
	apiVersion.POST("/login", auth.Login)
}
