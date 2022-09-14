package main

import "fmt"

func emptyFunc(i int) interface{} {
	if i == 1 {
		return true
	} else if i == 2 {
		return "return string"
	} else {
		return 1
	}
	// return 1234
}

func main() {
	empty := emptyFunc(2)
	fmt.Println(empty)
}
