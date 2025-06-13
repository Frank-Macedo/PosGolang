package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Classe struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type Product struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	ClasseId     int
	Classe       Classe
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Product{}, &Category{}, &Classe{}, SerialNumber{})

	// db.Create(&Product{Name: "Notebook", Price: 1899.90})

	// products := []Product{
	// 	{Name: "Notebook", Price: 1899.90},
	// 	{Name: "Mouse de ouro", Price: 5999.90},
	// 	{Name: "fone sem fio xing ling", Price: 29.90},
	// }

	// db.Create(&products)

	// var product Product
	// db.First(&product, "name = ?", "Notebook")
	// fmt.Printf("product: %s, possui o preço de %.2f\n", product.Name, product.Price)

	// var product []Product

	// db.Find(&product, "Price > ?", 1800)

	// for _, p := range product {
	// 	fmt.Printf("product: %s, possui o preço de %.2f\n", p.Name, p.Price)
	// }

	// var p Product

	// db.First(&p, 1)
	// fmt.Printf("product: %s, possui o preço de %.2f\n", p.Name, p.Price)

	// p.Name = "Notebook gamer usado"

	// db.Save(&p)

	// db.First(&p, 1)
	// fmt.Printf("product: %s, possui o preço de %.2f\n", p.Name, p.Price)

	classe := Classe{Name: "Normal"}
	db.Create(&classe)

	category := Category{Name: "Cozina"}
	db.Create(&category)

	product := Product{Name: "Panela", Price: 1899.90, CategoryID: category.ID, ClasseId: classe.ID}
	db.Create(&product)

	SerialNumber := SerialNumber{Number: "123456789", ProductID: product.ID}
	db.Create(&SerialNumber)

	var products []Product
	db.Preload("Category").Preload("Classe").Preload("SerialNumber").Find(&products)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err.Error())
	}

	for _, c := range categories {
		fmt.Printf("Category: %s\n", c.Name)
		for _, p := range c.Products {
			fmt.Printf("Product: %s\n", p.Name)
		}
	}
}
