//go:build wireinject
// +build wireinject

package core

import "github.com/google/wire"

func InitializeBackendApp() (App, error) {
	wire.Build(DefaultBackendApplication)
	return nil, nil
}

func InitializeAsyncWorker() (App, error) {
	wire.Build(DefaultBackendAsyncWorker)
	return nil, nil
}
