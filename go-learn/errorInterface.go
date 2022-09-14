package main

import (
	"errors"
	"fmt"
)

type Error interface {
	Error() string
}

func DividedBy(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi dengan nol N/A")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	bagibagian, error := DividedBy(3, 0)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println(bagibagian)
	}

}
