package controller

import (
	"TapTalk-BE/auth"
	"TapTalk-BE/database"
	"TapTalk-BE/model/User"
	"encoding/json"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(username, password string) (interface{}, error) {
	db := database.ConnDb()
	var err error
	user := User.User{}

	emails := db.Debug().Model(User.User{}).Where("email = ?", username).Take(&user).RowsAffected
	if emails == 0 {
		err = db.Debug().Model(User.User{}).Where("username = ?", username).Take(&user).Error
		if err != nil {
			return "", err
		}
		err = User.VerifyPassword(user.Password, password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return "password wrong. Try Again", http.ErrBodyNotAllowed
		}
		return auth.CreateToken(user.Id)
	}
	err = User.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "password wrong. Try Again", http.ErrBodyNotAllowed
	}
	return auth.CreateToken(user.Id)
}

func Login(c echo.Context) error {
	user := User.Login{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cred, err := SignIn(user.Username, user.Password)
	if err != nil {
		return c.String(http.StatusForbidden, "Wrong Username/Password")
	}

	return c.JSON(http.StatusOK, cred)
}
