package ports

import "github.com/ThailanTec/rascunho/internal/core/domain"

type UserRepository interface {
	CreateUser(user domain.User) (u *domain.User, e error)
	GetAllUser() (u *domain.User, e error)
	GetUser(nome string) (u *domain.User, e error)
	DeleteUser(id int) (delete bool, e error)
}
