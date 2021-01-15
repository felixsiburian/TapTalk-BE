package controller

import (
	"TapTalk-BE/database"
	"TapTalk-BE/model/User"
	"encoding/json"
	"github.com/labstack/echo"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

func isPasswordValid(s string) bool {
	var (
		isMinLen = false
		isUpper  = false
		isLower  = false
		isNumber = false
		isSymbol = false
	)
	if len(s) >= 6 && len(s) <= 32 {
		isMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			isUpper = true
		case unicode.IsLower(char):
			isLower = true
		case unicode.IsNumber(char):
			isNumber = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			isSymbol = true
		}
	}
	return isMinLen && isUpper && isLower && isNumber && isSymbol
}

func CheckCredentialExist(email, username string) bool {
	db := database.ConnDb()
	user := User.User{}
	emails := db.Debug().Model(&User.User{}).Where("email = ?", email).Take(&user).RowsAffected
	if emails > 0 {
		return true
	}
	usernames := db.Debug().Model(&User.User{}).Where("username = ?", username).Take(&user).RowsAffected
	if usernames > 0 {
		return true
	}
	return false
}

func Register(c echo.Context) error {
	db := database.ConnDb()
	user := new(User.User)

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	checkCredential := CheckCredentialExist(user.Email, user.Username)
	if checkCredential == true {
		return c.String(http.StatusForbidden, "Email or Username already Exist")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if len(user.Fullname) <= 0 {
		return c.String(http.StatusBadRequest, "Invalid Fullname")
	}
	pass := isPasswordValid(user.Password)
	if pass != true {
		return c.String(http.StatusBadRequest, "Password is not feasible")
	}
	if len(user.Password) < 6 && len(user.Password) > 32 {
		return c.String(http.StatusBadRequest, "Invalid Password")
	}
	emails := isEmailValid(user.Email)
	if emails != true {
		return c.String(http.StatusBadRequest, "Email is not feasible")
	}
	if len(user.Email) <= 0 {
		return c.String(http.StatusBadRequest, "Invalid Email")
	}
	if len(user.Birthday.String()) <= 0 {
		return c.String(http.StatusBadRequest, "Invalid birth date")
	}
	if len(user.Username) <= 0 {
		return c.String(http.StatusBadRequest, "Invalid username")
	}
	hashedPaswword, err := User.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user.Password = string(hashedPaswword)
	currentTime := time.Now()
	user.CreatedDate = currentTime.Format(formatDate)
	err = db.Debug().Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
