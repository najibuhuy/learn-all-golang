package main

import "fmt"

func logging() {
	fmt.Println("Selesai run Application")
}
func runApplication(number int) {
	defer logging() //after runApplication() is done, logging() running whatever error or not
	result := 10 / number
	fmt.Println("Result: ", result)
	fmt.Println("ini Run Application")
}

func main() {
	runApplication(0)
}
