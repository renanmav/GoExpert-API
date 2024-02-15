package entity

import (
	"github.com/renanmav/GoExpert-API/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	productName  = "Test Product"
	productPrice = 100.0
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
	assert.Error(t, product.ValidateProduct(), ErrNameIsRequired)

	product.Name = productName
	product.Price = -1 // Invalid product price
	assert.Error(t, product.ValidateProduct(), ErrInvalidPrice)
}

func TestProduct_ValidateProduct_IDIsRequired(t *testing.T) {
	product, _ := NewProduct(productName, productPrice)
	product.ID = entity.ID{}

	err := product.ValidateProduct()

	assert.NotNil(t, err)
	assert.Equal(t, ErrIDIsRequired, err)
}

func TestProduct_ValidateProduct_NameIsRequired(t *testing.T) {
	product, _ := NewProduct(productName, productPrice)
	product.Name = ""

	err := product.ValidateProduct()

	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_ValidateProduct_PriceIsRequired(t *testing.T) {
	product, _ := NewProduct(productName, productPrice)
	product.Price = 0

	err := product.ValidateProduct()

	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_ValidateProduct_InvalidPrice(t *testing.T) {
	product, _ := NewProduct(productName, productPrice)
	product.Price = -1

	err := product.ValidateProduct()

	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
