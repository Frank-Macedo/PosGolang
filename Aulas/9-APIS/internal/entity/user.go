package entity

import (
	"errors"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"string"`
	Name     string    `json:"name"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

var ErrEmptyPassword = errors.New("password cannot be empty")

func NewUser(name, password, email string) (*User, error) {

	if password == "" {
		return nil, ErrEmptyPassword
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
