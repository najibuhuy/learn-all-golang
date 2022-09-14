package main

import "fmt"

func main() {
	type noNIK string
	type Married bool

	var noNikNajib noNIK = "3219031313232132"
	fmt.Println(noNikNajib)
	fmt.Println(noNIK("2222222222"))

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
