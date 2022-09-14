package main

import "fmt"

func main() {
	const lastName = "alyasyfi" //can,t be changed

	fmt.Println(lastName)

	const (
		country   = "Germany"
		selection = "ausbildung"
	)

	fmt.Println("selection: ", selection)
	const testString = country + " " + selection
	fmt.Println(testString)

}
