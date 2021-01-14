package middleware

import (
	"TapTalk-BE/auth"
	"github.com/labstack/echo"
	"net/http"
)

func SetAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth.TokenValid(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}
