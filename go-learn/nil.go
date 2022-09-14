package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	ExMap := NewMap("najib")
	if ExMap == nil {
		fmt.Println("Hasil Nil")
	} else {
		fmt.Println(ExMap)
	}
}
