package main

// import(
// 	"fmt"
// 	"go-learn/sayhello"
// )

import (
	"fmt"
	"go-learn/helper"
)

func main() {
	result := helper.SayHello("Najib")
	resultHi := helper.SayHi("Nauval")
	fmt.Println(result)
	fmt.Println(resultHi)
	fmt.Println(helper.Application)
	// fmt.Println(helper.letter) //error

}
