package main

import "fmt"

func main() {

	person := map[string]int{
		"age":    10,
		"height": 120,
	}

	person["width"] = 800
	fmt.Println(person)

	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "najib"

	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)

}
