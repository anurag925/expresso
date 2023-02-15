//go:build wireinject
// +build wireinject

package repositories

import (
	"expresso/configs/core"
	"expresso/pkg/repositories/models"

	"github.com/google/wire"
)

func InitUserRepo(app core.App, user *models.User) *User {
	wire.Build(New, NewUserRepository)
	return &User{}
}

func InitUser(app core.App) *User {
	wire.Build(New, NewUser)
	return &User{}
}
