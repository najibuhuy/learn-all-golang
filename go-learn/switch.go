package main

import "fmt"

func main() {
	name := "abdul"
	fmt.Println()
	switch name {
	case "najib":
		fmt.Println("ni anak ganteng")
	case "none":
		fmt.Println("ni anak belagu")
	default:
		fmt.Println("ga ganteng")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama panjang")
	default:
		fmt.Println("nama pendek")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah bagus")

	}
}
