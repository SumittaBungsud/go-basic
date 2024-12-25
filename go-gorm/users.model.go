package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func createUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func loginUser(db *gorm.DB, user *User) (string,error) {
	selectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// create JWT token to return if the password is matched
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": selectedUser.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
