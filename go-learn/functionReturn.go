package main

import "fmt"

func bilangHello(name string) string {
	return "Hello " + name
}

func main() {
	name := bilangHello("ali")
	fmt.Println(name)
}
