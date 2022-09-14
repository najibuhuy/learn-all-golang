package main

import "fmt"

func main() {
	var nilai1 = 200
	var nilai2 = 300

	var lulusNilai = 250
	var lulusNaik = 200

	var lulusMinimal1 = nilai1 > lulusNaik && nilai1 > lulusNilai
	var lulusMinimal2 = nilai2 > lulusNaik && nilai2 > lulusNilai

	fmt.Println(lulusMinimal1)
	fmt.Println(lulusMinimal2)

}
