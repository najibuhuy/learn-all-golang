package main

import "fmt"

func main() {
	var bil1 = 9
	var bil2 = 10

	var div float32 = float32(bil1) / float32(bil2)
	fmt.Println("plus: ", bil1+bil2)
	fmt.Println("times: ", bil1*bil2)
	fmt.Println("divide: ", div)
	fmt.Println("minus: ", bil1-bil2)
	fmt.Println("modulus: ", bil1%bil2)

	//augmented assigment
	bil1 += 10 //bil1 = bil1 + 10
	fmt.Println(bil1)

	bil1 -= 15 //bil1 = bil1 - 15
	fmt.Println(bil1)

	//unary operator
	bil1++ //bil1 = bil1 + 1
	fmt.Println(bil1)

	bil1-- //bil1 = bil1 -1
	fmt.Println(bil1)

}
