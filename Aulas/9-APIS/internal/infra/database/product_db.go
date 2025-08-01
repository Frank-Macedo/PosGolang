package database

import (
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {
	if err := p.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error

	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {

	_, err := p.FindByID(uuid.UUID(product.ID).String())
	if err != nil {
		return err
	}
	if err := p.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	if err := p.DB.Delete(&entity.Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
