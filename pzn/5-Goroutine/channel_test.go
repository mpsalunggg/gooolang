package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		for i := 1; i <= 2; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("Detik :", i)
		}
		time.Sleep(2 * time.Second)
		channel <- "Hello"

		fmt.Println("Selesai mengirim pesan ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Detik :", i)
	}
}	