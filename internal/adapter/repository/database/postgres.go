package repository

import (
	"errors"
	"fmt"

	"github.com/ThailanTec/rascunho/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func NewMessengerPostgresRepository() *MessengerPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "pass1234"
	dbname := "postgres"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.User{})

	return &MessengerPostgresRepository{
		db: db,
	}
}

func (p *MessengerPostgresRepository) CreateUser(user domain.User) (u *domain.User, e error) {

	req := p.db.Create(&user)

	if req.RowsAffected == 0 {
		return nil, errors.New("erro ao criar usuario")
	}
	return u, e
}

func (p *MessengerPostgresRepository) GetUser(name string) (user *domain.User, e error) {
	user = &domain.User{}
	req := p.db.First(&user, "name = ? ", name)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (p *MessengerPostgresRepository) GetAllUser() (user *domain.User, e error) {
	var use []*domain.User
	req := p.db.Find(&use)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return user, nil
}

func (p *MessengerPostgresRepository) DeleteUser(id int) (delete bool, e error) {
	return
}
