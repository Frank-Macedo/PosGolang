package entity

import (
	"errors"
	"time"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
}


var (
	errIdsRequired   = errors.New("product ID is required")
	errNameRequired  = errors.New("product name is required")
	errPriceRequired = errors.New("product price is required")
	errInvalidPrice  = errors.New("product price must be greater than zero")
	errInvalidID     = errors.New("product ID is invalid")
)

func NewProduct(name string, price float64) (*Product, error) {

	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {

	// if _, err := entity.ParseID(string(p.ID)); err != nil {
	// 	return errInvalidID
	// }
	if p == nil {
		return errors.New("product cannot be nil")
	}
	if p.Name == "" {
		return errNameRequired
	}
	if p.Price <= 0 {
		return errInvalidPrice
	}
	return nil
}
