package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func PrintHello(name string) {
	fmt.Println("Hello", name)
}

func TestGoroutine(t *testing.T) {
	go PrintHello("World")
	time.Sleep(5 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number :", number)
}

func TestManyRandomNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}