package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Test Product", 10.0)
	assert.Nil(t, err, "Expected no error when creating a new product")
	assert.NotNil(t, p, "Expected product to be created")
	assert.NotEmpty(t, p.ID, "Expected product ID to be generated")
	assert.Equal(t, "Test Product", p.Name, "Expected product name to be 'Test Product'")
	assert.Equal(t, 10.0, p.Price, "Expected product price to be 10.0")
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p, "Expected product to be nil when name is empty")
	assert.Equal(t, errNameRequired, err, "Expected error to be 'product name is required'")
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Test Product", 0)
	assert.Nil(t, p, "Expected product to be nil when price is zero")
	assert.Equal(t, errInvalidPrice, err, "Expected error to be 'product price must be greater than zero'")
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Test Product", 10.0)
	assert.Nil(t, err, "Expected no error when creating a new product")
	assert.NotNil(t, p, "Expected product to be created")
	assert.Nil(t, p.Validate(), "Expected product validation to succeed")
}
