package main

import "fmt"

type Customer struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	var newCustomer Customer
	newCustomer.Age = 26
	newCustomer.FirstName = "najib"
	newCustomer.LastName = "alyasyfi"
	fmt.Println(newCustomer)

	//Struct Literals
	secondCustomer := Customer{
		FirstName: "Ali",
		LastName:  "Muhammad",
		Age:       22,
	}

	thirdCustomer := Customer{"Mehdi", "Viki", 18}

	fmt.Println(secondCustomer)
	fmt.Println(thirdCustomer)

}
