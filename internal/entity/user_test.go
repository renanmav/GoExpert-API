package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	userName     = "Renan"
	userEmail    = "me@renanmav.xyz"
	userPassword = "userPassword"
)

func TestUser_NewUser(t *testing.T) {
	user, err := NewUser(userName, userEmail, userPassword)

	assert.NoError(t, err)
	assert.Equal(t, user.Name, userName)
	assert.Equal(t, user.Email, userEmail)
	assert.NotEqual(t, user.Password, userPassword) // Password should be hashed, not plain text
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser(userName, userEmail, userPassword)

	// Should return true for correct userPassword
	assert.True(t, user.ValidatePassword(userPassword))

	// Should return false for incorrect userPassword
	assert.False(t, user.ValidatePassword("lol"))
}
