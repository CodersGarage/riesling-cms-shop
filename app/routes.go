package app

import "github.com/gin-gonic/gin"

func InitRoutes() {
	config := GetAppConfig()
	gin.SetMode(config.Get("default.app_mode").String()) // Mode of Application [debug, release]

	// Routes of Application Starts Here
	routes := gin.Default()

	// Routes of Version 1
	apiV1 := routes.Group("/api/v1")

	// Running Application on Configured Address host:port
	routes.Run(config.Get("default.app_uri").String())
}
