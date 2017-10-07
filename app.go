package main

import (
	"riesling-cms-core/app"
)

func main() {
	app.InitAppConfig()
	config := app.GetAppConfig()
	println("AppName : ", config.Get("default.appName").String())
	println("AppVersion : ", config.Get("default.appVersion").String())
	println("AppDeveloper : ", config.Get("default.developer").String())
}
