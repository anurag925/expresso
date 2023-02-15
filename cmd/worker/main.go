package main

import (
	"expresso/configs/core"
	"expresso/pkg/tasks"
	"fmt"
)

func main() {
	app, err := core.InitializeAsyncWorker()
	if err != nil {
		fmt.Printf("unable to initialize the application with error %+v", err)
	}
	tasks.New(app).RegisterTask()
	if err = app.Tasks().StartServer(); err != nil {
		fmt.Printf("unable to start the async task with error %+v", err)
	}
}
