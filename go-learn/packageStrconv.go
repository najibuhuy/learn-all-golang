package main

import (
	"fmt"
	"strconv"
)

func main() {
	exStr := "true"
	resultStr, err := strconv.ParseBool(exStr)
	if err == nil {
		fmt.Println(resultStr)
	} else {
		fmt.Println(err.Error())
	}

	number, errs := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(errs.Error())
	}
}
