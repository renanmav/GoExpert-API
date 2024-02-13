package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	productName  = "Test Product"
	productPrice = 100
)

func TestProduct_NewProduct(t *testing.T) {
	product, err := NewProduct(productName, productPrice)

	assert.NoError(t, err)
	assert.Equal(t, product.Name, productName)
	assert.Equal(t, product.Price, productPrice)
}

func TestProduct_ValidateProduct(t *testing.T) {
	product, _ := NewProduct(productName, productPrice)

	assert.NoError(t, product.ValidateProduct())

	product.Name = "" // Invalid product name
	assert.Error(t, product.ValidateProduct())

	product.Name = productName
	product.Price = -1 // Invalid product price
	assert.Error(t, product.ValidateProduct())
}
