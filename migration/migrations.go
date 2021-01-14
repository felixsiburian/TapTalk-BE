package migration

import (
	"TapTalk-BE/database"
	"TapTalk-BE/model/User"
	"fmt"
)

func Migrate() {
	db := database.ConnDb()
	db.AutoMigrate(&User.User{})
	fmt.Println("Succesfully Migrate")
}
