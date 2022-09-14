package main

import "fmt"

func namaNama(nama string, age int) (string, int) {
	return nama, age
}

func main() {
	name, age := namaNama("najib", 26)
	fmt.Println(namaNama("ali", 20))
	fmt.Println(age, name)
}
