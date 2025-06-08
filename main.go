package main

import (
	"dopc-service/routers"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Load routes
	routers.InitializeRoutes()

	// Access app.conf settings
	port, _ := web.AppConfig.Int("httpport")
	appName, _ := web.AppConfig.String("appname")

	// Set up logging
	logs.SetLogger(logs.AdapterConsole)
	fmt.Printf("Starting %s on port %d...\n", appName, port)

	// Start the application
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
