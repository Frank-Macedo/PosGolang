package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Frank", "123456", "frank@321.com")

	assert.Nil(t, err, "Expected no error when creating a new user")
	assert.NotNil(t, user, "Expected user to be created")
	assert.NotEmpty(t, user.ID, "Expected user ID to be generated")
	assert.NotEmpty(t, user.Password, "Expected user password to be hashed")
	assert.Equal(t, "Frank", user.Name, "Expected user name to be 'Frank'")
	assert.Equal(t, "frank@321.com", user.Email, "Expected user email to be")

}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Frank", "123456", "frank@321.com")
	assert.Nil(t, err, "Expected no error when creating a new user")
	assert.True(t, user.ValidatePassword("123456"), "Expected password validation to succeed with correct password")
	assert.False(t, user.ValidatePassword("wrongpassword"), "Expected password validation to fail with incorrect password")
}
