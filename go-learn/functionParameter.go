package main

import "fmt"

type Filter func(string) string

func sayHelloDong(name string, filterKataKasar Filter) string {
	filter := filterKataKasar(name)
	return "hello " + filter
}

func filterKataKasar(message string) string {
	if message == "anjing" {
		return "..."
	} else {
		return message
	}
}

func main() {
	filters := filterKataKasar
	hellos := sayHelloDong("anjing", filters)
	fmt.Println(hellos)
}
