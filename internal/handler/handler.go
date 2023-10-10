// internal/handler/handler.go
package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupHandlers(
	apiHandler *ApiHandler,
) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	apiRouterGroup := router.Group("/current")
	apiRouterGroup.GET("/", apiHandler.List)

	return router
}
