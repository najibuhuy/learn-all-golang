package main

import "fmt"

type Customer struct {
	FirstName, LastName string
	Age                 int
}

func (customer Customer) sayHello(name string) string {
	return "Hello " + customer.FirstName + " " + customer.LastName + " My name is " + name
}
func main() {
	najib := Customer{
		FirstName: "najib",
		LastName:  "Alyasyfi",
	}
	coba := najib.sayHello("gogon")
	fmt.Println(coba)
}
