package main

import (
	"TapTalk-BE/router"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome ...")
	e := router.New()
	err := e.Start(":8080")
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
