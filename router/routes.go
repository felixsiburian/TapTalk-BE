package router

import (
	"TapTalk-BE/controller"
	"TapTalk-BE/migration"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	migration.Migrate()
	e.POST("/Register", controller.Register)

	return e
}
