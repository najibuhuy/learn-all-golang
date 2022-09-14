package main

import (
	"fmt"
	_ "fmt"
	"go-learn/database"
	_ "go-learn/database" //blank identifier
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
