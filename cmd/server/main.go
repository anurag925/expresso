package main

import (
	"expresso/configs/core"
	"expresso/pkg/controllers"
	"expresso/pkg/tasks"
	"fmt"
)

func main() {
	app, err := core.InitializeBackendApp()
	if err != nil {
		fmt.Printf("unable to initialize the application with error %+v", err)
		return
	}
	routes := controllers.NewRoutes(app)
	routes.SetupMiddleware()
	routes.SetupRoutes()
	tasks.New(app).ScheduleTask()
	go app.Tasks().StartScheduler()
	app.Server().Start(app.Config().App().Port)
}
