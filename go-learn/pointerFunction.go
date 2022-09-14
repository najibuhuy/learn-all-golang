package main

import "fmt"

type DetailAddress struct {
	City, Province, Country string
}

func ChangeAddressToIndo(address *DetailAddress) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndoNoPointer(address DetailAddress) {
	address.Country = "Indonesia"
}

func main() {
	address := DetailAddress{
		"Bandung",
		"Jawa Barat",
		"",
	}
	address1 := DetailAddress{
		"Bogoro", "", "",
	}
	ChangeAddressToIndo(&address)
	ChangeAddressToIndoNoPointer(address1)
	fmt.Println(address)
	fmt.Println(address1) // it will not change

}
