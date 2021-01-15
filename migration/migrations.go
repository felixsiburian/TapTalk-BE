package migration

import (
	"TapTalk-BE/database"
	"TapTalk-BE/model/DailyEntries"
	"TapTalk-BE/model/User"
	"fmt"
)

func Migrate() {
	db := database.ConnDb()
	db.AutoMigrate(&User.User{}, &DailyEntries.DailyEntries{})
	fmt.Println("Succesfully Migrate")
}
