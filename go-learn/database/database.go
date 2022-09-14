package database

import "fmt"

var connection string

func init() {
	connection = "no sql"
	fmt.Println("init ke panggil")
}

func GetDatabase() string {
	return connection
}
