package product

import "database/sql"

type ProductRepositoryTxt struct {
	db *sql.DB
}

func NewProductRepositoryTxt(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepositoryTxt{}
}

func (r *ProductRepositoryTxt) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product Name TXT",
	}, nil
}
