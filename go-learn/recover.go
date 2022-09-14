package main

import (
	"fmt"
)

func endApp() {
	message := recover() //catch data error
	if message != nil {
		fmt.Println("Error Message: ", message)
	}
	fmt.Println("End Application")
}

/*
komentar bro
*/

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error uy") //stop the script in this function
	}
}

func main() {
	runApp(false)
}
