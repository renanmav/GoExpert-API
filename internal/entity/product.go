package entity

import (
	"errors"
	"github.com/renanmav/GoExpert-API/pkg/entity"
	"time"
)

// Product errors
var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
}

// ValidateProduct validates the product
func (p *Product) ValidateProduct() error {
	var emptyID = entity.ID{}
	if p.ID.String() == emptyID.String() || p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}

// NewProduct creates a new product
func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().String(),
	}
	if err := p.ValidateProduct(); err != nil {
		return nil, err
	}
	return p, nil
}
