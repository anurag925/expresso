package controllers

import "expresso/configs/core"

// Controller provides base functionality and dependencies to routes.
// The proposed pattern is to embed a Controller in each individual route struct and to use
// the router to inject the container so your routes have access to the services within the container
type Controller struct {
	// Container stores a services container which contains dependencies
	core.App
}

// NewController creates a new Controller
func NewController(a core.App) *Controller {
	return &Controller{
		App: a,
	}
}
