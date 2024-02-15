package database

import (
	"fmt"
	"github.com/renanmav/GoExpert-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestProductDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Cube", 10)
	if err != nil {
		t.Error(err)
	}

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProductDB_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProductDB(db)

	// Create a product
	product, err := entity.NewProduct("Test Product", 10)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	if err != nil {
		t.Error(err)
	}

	// Test FindByID
	foundProduct, err := productDB.FindByID(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, product.ID, foundProduct.ID)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
}

func TestProductDB_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProductDB(db)

	// Create a product
	product, err := entity.NewProduct("Test Product", 10.0)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	if err != nil {
		t.Error(err)
	}

	// Update the product
	product.Name = "Updated Product"
	product.Price = 20.0

	// Test Update
	err = productDB.Update(product)
	if err != nil {
		t.Error(err)
	}

	// Find the updated product
	updatedProduct, err := productDB.FindByID(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, product.ID, updatedProduct.ID)
	assert.Equal(t, "Updated Product", updatedProduct.Name)
	assert.Equal(t, 20.0, updatedProduct.Price)
}

func TestProductDB_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProductDB(db)

	// Create a product
	product, err := entity.NewProduct("Test Product", 10)
	if err != nil {
		t.Error(err)
	}

	err = productDB.Create(product)
	if err != nil {
		t.Error(err)
	}

	// Test Delete
	err = productDB.Delete(product.ID.String())
	if err != nil {
		t.Error(err)
	}

	// Try to find the deleted product
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err, "record not found")
}

func TestProductDB_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProductDB(db)

	// Create a few products
	for i := 0; i < 5; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), 10)
		if err != nil {
			t.Error(err)
		}

		err = productDB.Create(product)
		if err != nil {
			t.Error(err)
		}
	}

	// Test FindAll
	products, err := productDB.FindAll(1, 3, "asc")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, len(products))
	assert.Equal(t, "Product 0", products[0].Name)
	assert.Equal(t, "Product 1", products[1].Name)
	assert.Equal(t, "Product 2", products[2].Name)

	products, err = productDB.FindAll(2, 3, "asc")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(products))
	assert.Equal(t, "Product 3", products[0].Name)
	assert.Equal(t, "Product 4", products[1].Name)
}
