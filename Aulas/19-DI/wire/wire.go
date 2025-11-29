//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/Frank-Macedo/19-DI/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
)

func InitializeProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(setRepositoryDependency, product.NewProductUseCase)
	return &product.ProductUseCase{}
}
