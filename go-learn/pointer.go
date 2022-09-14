package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		"Bandung",
		"Jawa barat",
		"Indonesia",
	} //struct
	address2 := address1
	address2.City = "Bogor"
	fmt.Println(address1) // address1 will not not change
	fmt.Println(address2)

	//pointer > pass by reference, when we change address2 value, address1 will be changed too
	address3 := &address1
	address3.Country = "Nigeria"
	address3.City = "Bogoro"
	fmt.Println(address3, "address3")
	fmt.Println(address1, "address1")
	fmt.Println(address2, "address2")
	fmt.Println(address3.Country, "address3.Country")

	//pointer make new data
	address4 := &address1
	address4 = &Address{
		"Jakarta", "DKI Jakarta", "Indonesia",
	}
	fmt.Println(address4, "address4")
	fmt.Println(address1, "address1")

	//address 1 pointer to address 5
	address5 := &address1

	*address5 = Address{
		"Malang", "Jawa Timur", "Indonesia",
	}
	fmt.Println(address5, "address5")
	fmt.Println(address1, "address1")

	//new
	address6 := new(Address)
	address6.City = "Sumbing"
	fmt.Println(address6)

}
