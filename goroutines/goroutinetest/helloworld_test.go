package goroutinetest

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutines(t *testing.T) {
	go Helloworld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
