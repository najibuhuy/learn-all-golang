package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai endApp")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR") //stop the running script unless defer
	}
	fmt.Println("APlikasi berjalan")
}

func main() {
	runApp(true)
}
