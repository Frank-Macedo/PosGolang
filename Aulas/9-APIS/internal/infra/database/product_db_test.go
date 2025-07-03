package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer cleanupDatabase(db)

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Test Product", 100.0)

	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err, "Expected no error when creating product")
	assert.NotEmpty(t, product.ID, "Expected product ID to be generated")
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer cleanupDatabase(db)

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	// Create some products
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)

		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}
		err = productDB.Create(product)
		assert.NoError(t, err, "Expected no error when creating product")
	}

	// Find all products
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err, "Expected no error when finding all products")
	assert.Len(t, products, 10, "Expected to find 10 products")
	assert.Equal(t, "Product 1", products[0].Name, "Expected first product to be 'Product 1'")
	assert.Equal(t, "Product 10", products[9].Name, "Expected first product to be 'Product 10'")

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err, "Expected no error when finding all products")
	assert.Len(t, products, 10, "Expected to find 10 products")
	assert.Equal(t, "Product 11", products[0].Name, "Expected first product to be 'Product 11'")
	assert.Equal(t, "Product 20", products[9].Name, "Expected first product to be 'Product 20'")

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err, "Expected no error when finding all products")
	assert.Len(t, products, 3, "Expected to find 3 products")
	assert.Equal(t, "Product 21", products[0].Name, "Expected first product to be 'Product 21'")
	assert.Equal(t, "Product 23", products[2].Name, "Expected first product to be 'Product 23'")

}

func TestFindProductByID(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Product{})
	product := NewProduct(db)
	p, err := entity.NewProduct("Teste Produto", 100.00)
	if err != nil {
		fmt.Println("Error ocurred when create a Product")
	}
	err = product.Create(p)
	assert.NoError(t, err)
	id, _ := p.ID.Value()
	recuperedP, err := product.FindByID(id.(string))
	assert.NoError(t, err)
	assert.Equal(t, "Teste Produto", recuperedP.Name)
	assert.Equal(t, 100.00, recuperedP.Price)
}

func cleanupDatabase(db *gorm.DB) {
	if err := db.Exec("DELETE FROM products").Error; err != nil {
		fmt.Printf("Failed to clean up database: %v\n", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer cleanupDatabase(db)

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	// Create a product
	product, err := entity.NewProduct("Test Product", 100.0)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err, "Expected no error when creating product")

	// Update the product
	product.Name = "Updated Product"
	product.Price = 150.0
	err = productDB.Update(product)
	assert.NoError(t, err, "Expected no error when updating product")

	// Find the updated product
	updatedProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err, "Expected no error when finding updated product")
	assert.Equal(t, "Updated Product", updatedProduct.Name, "Expected updated product name to match")
	assert.Equal(t, 150.0, updatedProduct.Price, "Expected updated product price to match")
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer cleanupDatabase(db)

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	// Create a product
	product, err := entity.NewProduct("Test Product", 100.0)
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err, "Expected no error when creating product")

	// Delete the product
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err, "Expected no error when deleting product")

	// Try to find the deleted product
	deletedProduct, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err, "Expected error when finding deleted product")
	assert.Nil(t, deletedProduct, "Expected deleted product to be nil")
}
