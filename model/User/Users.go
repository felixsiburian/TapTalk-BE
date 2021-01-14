package User

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id          uint32    `json:"user_id" gorm:"primary_key;auto_increment"`
	Fullname    string    `json:"fullname"`
	Birthday    time.Time `json:"birthday"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"created_date"`
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
