// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package services

import (
	"context"
	"expresso/configs/core"
)

// Injectors from wire.go:

func InitStatusService(app core.App, ctx context.Context) *StatusService {
	service := New(app, ctx)
	statusService := NewStatusService(service)
	return statusService
}
