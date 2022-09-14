package goroutinetest

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Najib"
		fmt.Println("selesai mengirim data ke chnnel")

	}()

	data := <-channel
	fmt.Println(data)
	// time.Sleep(2 * time.Second)

}
