package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your db host") //-host=localhost, defaultValue, realValue
	username := flag.String("username", "root", "Put your db user")
	password := flag.String("password", "root", "Put your db password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
