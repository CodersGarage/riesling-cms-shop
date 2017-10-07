package app

import "github.com/gin-gonic/gin"

func InitRoutes() {
	config := GetAppConfig()
	gin.SetMode(config.Get("default.appMode").String()) // Mode of Application [debug, release]

	// Routes of Application Starts Here
	routes := gin.Default()

	// Routes of Version 1
	apiV1 := routes.Group("/api/v1")
	apiV1.GET("/user/create")

	// Running Application on Configured Address host:port
	routes.Run(config.Get("default.appUri").String())
}
