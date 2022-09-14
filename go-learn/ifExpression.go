package main

import "fmt"

func main() {
	name := "najb"
	if name == "najib" {
		fmt.Println("hi najib")
	} else if name == "joke" {
		fmt.Println("halo joko")
	} else {
		fmt.Println("bukan najib")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama culkup")
	} else {
		fmt.Println("nama terlalu sedikit")
	}
}
