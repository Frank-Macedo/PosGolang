package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"Many2Many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	Price      float64
	Categories []Category `gorm:"Many2Many:product_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Cozina"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)

	// product := Product{Name: "Panela", Price: 1899.90, Categories: []Category{category, category2}}
	// db.Create(&product)

	// var products []Product
	// db.Preload("Category").Preload("Classe").Preload("SerialNumber").Find(&products)

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
