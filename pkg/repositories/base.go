package repositories

import "expresso/configs/core"

type Repository struct {
	core.App
}

func New(app core.App) *Repository {
	return &Repository{app}
}
