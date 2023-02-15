package repositories

import "expresso/pkg/repositories/models"

type User struct {
	*Repository
	*models.User
}

func NewUserRepository(repo *Repository, user *models.User) *User {
	return &User{repo, user}
}

func NewUser(repo *Repository) *User {
	return &User{repo, &models.User{}}
}
