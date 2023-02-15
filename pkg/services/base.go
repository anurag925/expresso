package services

import (
	"context"
	"expresso/configs/core"
	"expresso/configs/core/initializers"
)

type Service struct {
	app core.App
	ctx context.Context
}

func New(a core.App, ctx context.Context) *Service {
	return &Service{
		app: a,
		ctx: ctx,
	}
}

func (s *Service) Logger() initializers.Logger {
	return s.app.Logger()
}
