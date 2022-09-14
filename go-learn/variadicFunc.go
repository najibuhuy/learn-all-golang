package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers { // jadi slices
		total += number // total = total + number
	}
	return total
}

func main() {
	sum := sumAll(10, 20, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println(sum)

	mySlice := []int{20, 30, 40, 50}
	myresult := sumAll(mySlice...)
	fmt.Println(myresult)
}
