package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	randomBgt := random()
	// resultString := randomBgt.(string)
	// fmt.Println(resultString)

	// resultInt := randomBgt.(int) //panic
	// fmt.Println(resultInt)

	//use safe way

	switch result := randomBgt.(type) {
	case string:
		fmt.Println("string", result)
	case int:
		fmt.Println("int", result)
	default:
		fmt.Println("unknown")
	}
}
