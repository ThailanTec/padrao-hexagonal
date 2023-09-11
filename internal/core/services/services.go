package services

import (
	"github.com/ThailanTec/rascunho/internal/core/domain"
	"github.com/ThailanTec/rascunho/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(nuser domain.User) (user *domain.User, e error) {
	return u.repo.CreateUser(nuser)
}

func (u *UserService) GetAllUser() (user *domain.User, e error) {
	return u.repo.GetAllUser()
}

func (u *UserService) GetUser(name string) (user *domain.User, e error) {
	return u.repo.GetUser(name)
}

func (u *UserService) DeleteUser(id int) (delete bool, e error) {
	return u.repo.DeleteUser(id)
}
