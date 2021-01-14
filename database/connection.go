package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=TapTalk dbname=TapTalk-DB sslmode=disable password=taptalk123")
	if err != nil {
		fmt.Println("Erorr db : ", err.Error())
		panic("failed to connect database")
	}
	return db
}
