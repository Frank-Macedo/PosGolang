package main

import (
	"database/sql"
	"fmt"

	"github.com/Frank-Macedo/19-DI/wire"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err.Error())
	}
	productUseCase := wire.InitializeProductUseCase(db)
	product, _ := productUseCase.GetProduct(1)

	fmt.Println(product.Name)
}
