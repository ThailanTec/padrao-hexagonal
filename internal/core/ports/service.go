package ports

import "github.com/ThailanTec/rascunho/internal/core/domain"

type UserService interface {
	CreateUser(user domain.User) (domain.User, error)
	GetAllUser() (domain.User, error)
	GetUser(email string) (domain.User, error)
	DeleteUser(id int) (delete bool, e error)
}
