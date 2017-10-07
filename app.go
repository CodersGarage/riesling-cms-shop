package main

import (
	"riesling-cms-core/app"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		app.InitAppConfigWithPath(args[1])
	} else {
		app.InitAppConfig()
	}

	app.PrintLog("main", "Application configuration complete.")
	app.PrintLog("main", "Application Running {"+
		"\n\tAppMode : "+ app.GetAppConfig().Get("default.appMode").String()+
		"\n\tAppAddress : "+ app.GetAppConfig().Get("default.appUri").String()+
		"\n}")
	app.PrintLog("main", "Initializing Routes...")
	app.InitRoutes()
}
