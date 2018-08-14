package datamodels

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	 ID 		    int64     `json:"id" from:"id"`
	 Firstname      string    `json:"firstname" from:"firstname"`
	 Username       string    `json:"username" from:"username"`
	 HashedPassword []byte    `json:"-" from:"-"`
	 CreateAt       time.Time `json:"create_at" from:"created_at"`
}

func (u User) IsValid() bool {
	return u.ID > 0
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return  bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, nil
	}

	return true, nil
}