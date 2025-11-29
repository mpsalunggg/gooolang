package main

import (
	"fmt"
	"time"
)

func say(message string, repeat int) {
	for i := 1; i <= repeat; i++ {
		fmt.Printf("%d: %s\n", i, message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go say("Hello from Goroutine 1", 5)
	go say("Hello from Goroutine 2", 5)

	say("Main Function Running", 3)
	time.Sleep(2 * time.Second)
	fmt.Println("Program selesai.")
}
