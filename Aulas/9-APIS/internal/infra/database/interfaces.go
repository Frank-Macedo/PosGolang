package database

import (
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailId string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
