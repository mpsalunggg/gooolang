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