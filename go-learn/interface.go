package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(name HashName) string {
	return "Helo " + name.GetName()
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (names Animal) GetName() string {
	return names.Name
}

func main() {
	najib := Person{
		Name: "najib",
	}

	harimau := Animal{
		Name: "Harimau",
	}
	fmt.Println(SayHello(najib))
	fmt.Println(SayHello(harimau))
}
