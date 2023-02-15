package services

import (
	"expresso/pkg/repositories"
	"expresso/pkg/repositories/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type StatusService struct {
	*Service
}

func NewStatusService(s *Service) *StatusService {
	return &StatusService{s}
}

func (s *StatusService) Status() string {
	s.Logger().Info("logs showing up")
	return "OK"
}

func (s *StatusService) DBSetGetStatus() string {
	s.Logger().Info("logs showing up")
	userRepo := repositories.InitUser(s.app)
	userRepo.User.Name = null.StringFrom("hello")
	if err := userRepo.InsertG(boil.Infer()); err != nil {
		return err.Error()
	}
	if users, err := models.Users().AllG(); err != nil {
		return err.Error()
	} else {
		return users[0].Name.String
	}
	// return "OK"
}
