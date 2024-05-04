package service

import (
	"tube/pkg/config"
	"tube/pkg/model"
	"tube/pkg/repository"
)

type UserService interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository, config *config.Application) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetUserByID(id int64) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
