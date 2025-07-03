package database

import (
	"testing"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	user, _ := entity.NewUser("John Doe", "password123", "123456")
	userDB := NewUser(db)
	err = userDB.Create(user)

	assert.Nil(t, err, "Expected no error when creating user")

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err, "Expected no error when finding user by ID")
	assert.Equal(t, user.ID, userFound.ID, "Expected user ID to match")
	assert.Equal(t, user.Name, userFound.Name, "Expected user name to match")
	assert.Equal(t, user.Email, userFound.Email, "Expected user email to match")
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	user, _ := entity.NewUser("John Doe", "password123", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err, "Expected no error when creating user")

	foundUser, err := userDB.FindByEmail("123456")
	assert.Nil(t, err, "Expected no error when finding user by email")
	assert.Equal(t, user.ID, foundUser.ID, "Expected found user ID to match")
	assert.Equal(t, user.Name, foundUser.Name, "Expected found user name to match")
	assert.Equal(t, user.Email, foundUser.Email, "Expected found user email to match")
	assert.NotNil(t, foundUser.Password, "Expected found user to not be nil")

}
