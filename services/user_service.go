package services

import (
	"raytheon/datamodels"
	"raytheon/repository"
)

type UserService interface {
	GetAllUser() []datamodels.User
	GetUserByID(id int64) (datamodels.User, bool)
	DeleteUserByID(id int64) bool
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUser() []datamodels.User {
	panic("implement me")
}

func (s *userService) GetUserByID(id int64) (datamodels.User, bool) {
	panic("implement me")
}

func (s *userService) DeleteUserByID(id int64) bool {
	panic("implement me")
}
