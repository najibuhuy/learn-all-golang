package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "muhammad"
	name[1] = "najib"
	name[2] = "alyasyfi"
	fmt.Println(name[2])

	var values = [5]int{
		90, 80, 70,
	}

	fmt.Println(values[2])
	values[2] = 0

	fmt.Println(values[2])
	fmt.Println(len(values))
	fmt.Println(values)

}
