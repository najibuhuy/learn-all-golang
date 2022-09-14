package main

import "fmt"

func goodBye(name string) string {
	return "goodbye" + name
}

func main() {
	getGoodBye := goodBye
	fmt.Println(getGoodBye("najib"))
}
