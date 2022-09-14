package main

import "fmt"

func returnName() (age int, firstName, lastName string) {
	firstName = "muh najib"
	lastName = "alyasyfi"
	age = 26
	return
}

func main() {
	fName, lName, age := returnName()
	fmt.Println(fName, lName, age)
}
