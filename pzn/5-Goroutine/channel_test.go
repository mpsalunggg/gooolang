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

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello"
		fmt.Println("Selesai mengirim pesan ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	fmt.Println("Selesai mengirim pesan ke channel")

	time.Sleep(5 * time.Second)
}

// Channel as parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "mpsss"
	fmt.Println("Selesai mengirim pesan ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel in & out

func OnlyIn(channel chan<- string) {
	time.Sleep(4 * time.Second)
	channel <- "mpsss"
	fmt.Println("Selesai mengirim pesan ke channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
	fmt.Println("Selesai mengirim pesan ke channel")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Channel buffered
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Hello"
	channel <- "World"
	channel <- "Golang"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	go func() {
		fmt.Println("masuk")
		channel <- "Hello"
		channel <- "World"
		channel <- "Golang"
	}()

	time.Sleep(2 * time.Second)

	// go func() {
	// 	fmt.Println(<-channel)
	// 	fmt.Println(<-channel)
	// 	fmt.Println(<-channel)
	// 	fmt.Println("done")
	// }()

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// time.Sleep(4 * time.Second)
}
