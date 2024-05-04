package service

import (
	"tube/pkg/config"
	"tube/pkg/repository"
)

type Service struct {
	Auth AuthService
	User UserService
}

func NewService(repo repository.Repository, config *config.Application) *Service {
	return &Service{
		Auth: NewAuthService(repo, config),
		User: NewUserService(repo, config),
	}
}
