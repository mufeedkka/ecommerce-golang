package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Email       string    `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password    string    `json:"password" gorm:"not null" validate:"required"`
	Role        string    `json:"role" gorm:"not null;default:CLIENT"`
	BlockStatus bool      `json:"blockStatus" gorm:"not null;default:false"`
	Profile     Profile   `gorm:"foreignKey:UserId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Profile struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	UserId      int    `json:"userId" gorm:"foreignKey:UserId"`
	PhoneNumber int    `json:"phoneNumber"`
	Country     string `json:"country"`
	City        string `json:"city"`
	PinCode     int    `json:"pinCode"`
}

func (user *Users) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("cannot hash password. Try again")
	}

	user.Password = string(hash)
	return nil
}

func (user *Users) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("the provided password is wrong. Try again")
	}
	return nil
}
