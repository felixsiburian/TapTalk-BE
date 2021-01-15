package router

import (
	"TapTalk-BE/controller"
	"TapTalk-BE/middleware"
	"TapTalk-BE/migration"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	migration.Migrate()
	e.POST("/Register", controller.Register)
	e.POST("/Login", controller.Login)

	e.POST("/DailyEntries", controller.AddDailyEntries, middleware.SetAuth)
	e.GET("/GetDailyEntries/year=:year/quarter=:quarter", controller.GetDailyEntries)

	return e
}
