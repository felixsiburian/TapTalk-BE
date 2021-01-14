package auth

import (
	"TapTalk-BE/model/Token"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

func CreateToken(customer_id uint32) (Token.Token, error) {
	errs := os.Setenv("API_SECRET", "98hbun98h")
	if errs != nil {
		log.Fatal(errs.Error())
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["customer_id"] = customer_id

	//expired token
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		log.Fatal(err.Error())
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	stringRefToken, err := refreshToken.SignedString([]byte(os.Getenv("API_SECRET")))
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokens := Token.Token{
		Token:               stringToken,
		RefreshToken:        stringRefToken,
		TokenExpired:        time.Now().Add(time.Hour * 1).Unix(),
		RefreshTokenExpired: time.Now().Add(time.Hour * 24).Unix(),
	}

	return tokens, nil
}

func TokenValid(c echo.Context) error {
	errs := os.Setenv("API_SECRET", "98hbun98h")
	if errs != nil {
		log.Fatal(errs.Error())
	}
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(c echo.Context) string {
	keys := c.Request().URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request().Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenId(c echo.Context) (uint32, error) {
	errs := os.Setenv("API_SECRET", "98hbun98h")
	if errs != nil {
		return 0, errs
	}
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["customer_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}
