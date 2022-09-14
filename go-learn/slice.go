package main

import "fmt"

func main() {
	fmt.Println()
	var arrMonts = [12]string{
		"Jan",
		"Feb",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"Sept",
		"Octo",
		"Nov",
		"Des",
	}

	var sliceMonths = arrMonts[3:]
	var slicesMonths = arrMonts[10:11]
	fmt.Println(sliceMonths, "sliceMonts")
	fmt.Println(slicesMonths, "sliceMonts") //capacity = > [a:b] ===> a - len(array)

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"sunday",
	}
	daySlice1 := days[3:]
	daySlice1[0] = "liburrrr"

	fmt.Println(days)

	//append
	daySlice2 := append(daySlice1, "libur again")
	daySlice2[0] = "hari kebalikan"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "muh"
	newSlice[1] = "najib"

	fmt.Println(newSlice, "newSlice")
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
