package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("udah 5 udah")
			break //keluar for
		}
		fmt.Println(i)
	}

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue //ga keluar for, cuma stop pada perulangan ke genap
		}
		fmt.Println(i)
	}
}
