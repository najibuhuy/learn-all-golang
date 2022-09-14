package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fmt.Println()
	najib := &Man{
		"Najib",
	}

	najib.Married()
	fmt.Println(najib.Name)
}
