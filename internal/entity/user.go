package entity

import (
	"github.com/renanmav/GoExpert-API/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// User VO - Value Object - Entity - Model - DTO - Data Transfer Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"userName"`
	Email    string    `json:"userEmail"`
	Password string    `json:"-"`
}

// NewUser creates a new user and hashes the userPassword
func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password = string(hash)
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

// ValidatePassword compares the userPassword with the hash
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
