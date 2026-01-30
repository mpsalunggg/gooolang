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

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "Golang"
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(2 * time.Second)
}

// Select Channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "World"
	}()

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
