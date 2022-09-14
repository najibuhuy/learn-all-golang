package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args
	fmt.Println(argument)

	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println(hostName)
	} else {
		fmt.Println(err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
