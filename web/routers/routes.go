package routers

import (
	"fmt"
	"net/http"
	"raytheon/web/middleware/logger"
	"raytheon/web/middleware/requests"

	"raytheon/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(
		logger.Logger(utils.SetLogger("raytheon.log")),
		gin.Recovery(),
	)
	router.Use(requests.RequestIDMiddleware())
	router.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("X-Request-Id"))
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "url Not Found",
		})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "function not found",
		})
	})
	RegisterRouter(router)
	return router
}
