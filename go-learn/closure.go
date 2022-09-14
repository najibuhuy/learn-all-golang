package main

import "fmt"

// {

//} closure

func main() {
	counter := 0
	increment := func() {
		counter++
	}
	increment()

	fmt.Println(counter)
}
