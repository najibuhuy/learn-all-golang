package main

import "fmt"

func main() {
	counter := 0
	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	for counters := 0; counters <= 20; counters++ {
		fmt.Println("perulangan ke: ", counters)
	}

	newSlices := []string{
		"gondrong",
		"panjang",
		"cepak",
		"sorodot",
		"dukces",
	}

	for i, val := range newSlices {
		fmt.Println(i, val)
	}

}
