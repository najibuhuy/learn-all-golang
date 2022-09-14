package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackListFunc Blacklist) string {
	if blackListFunc(name) {
		return "You're blocked " + name
	} else {
		return "Welcome " + name
	}
}

func main() {
	blacklistFunc := func(name string) bool {
		if name == "najib" {
			return true
		} else {
			return false
		}
	} //anonymous func
	register := registerUser
	fmt.Println(register("admin", blacklistFunc))
	fmt.Println(register("najib", blacklistFunc))

}
