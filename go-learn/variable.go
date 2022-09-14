package main

import "fmt"

func main() {
	var name string

	name = "najib alyasyfi"
	fmt.Println(name)

	name = "muhammad najib"
	fmt.Println(name)

	var names = "najib budi"
	fmt.Println(names)

	var age = 30
	fmt.Println(age)

	country := "Indonesia" //only first declaration
	fmt.Println(country)

	var (
		firstName = "Najib"
		ages      = 31
	)

	fmt.Println(firstName)
	fmt.Println(ages)

	age = ages
	fmt.Println(age, "age")

}
