package main

import "fmt"

func main() {
	var name1 = "najib"
	var name2 = "najib"
	var name3 = "eko"

	var result1 bool = name1 == name2

	var result2 bool = name1 == name3

	fmt.Println(result1)
	fmt.Println(result2)

	var number1 = 100
	var number2 = 200

	var result3 = number1 > number2
	fmt.Println(result3)
}
