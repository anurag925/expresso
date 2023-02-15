//go:build wireinject
// +build wireinject

package services

import (
	"context"
	"expresso/configs/core"

	"github.com/google/wire"
)

func InitStatusService(app core.App, ctx context.Context) *StatusService {
	wire.Build(New, NewStatusService)
	return &StatusService{}
}
