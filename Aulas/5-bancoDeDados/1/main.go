package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price, ID: uuid.New().String()}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	product2 := NewProduct("Mouse de ouro", 5999.90)

	product3 := NewProduct("fone sem fio xing ling", 29.90)

	insertProduct(db, product)
	insertProduct(db, product2)
	insertProduct(db, product3)

	product.Name = "Notebook Gamer"
	product.Price = 2999.90
	err = changeProduct(db, product)
	if err != nil {
		panic(err.Error())
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("product: %s, possui o preço de %.2f\n", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err.Error())
	}

	for _, p := range products {
		fmt.Printf("product: %s, possui o preço de %.2f\n", p.Name, p.Price)
	}

	deleteProduct(db, product.ID)
	deleteProduct(db, product2.ID)
	deleteProduct(db, product3.ID)

}

func insertProduct(db *sql.DB, product *Product) error {

	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}

func changeProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
