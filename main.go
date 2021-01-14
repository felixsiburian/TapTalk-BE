package main

import (
	"TapTalk-BE/router"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome ...")
	e := router.New()
	err := e.Start("8000")
	if err != nil {
		log.Fatal("Error start : ", err.Error())
	}
}
