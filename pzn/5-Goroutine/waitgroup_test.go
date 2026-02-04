package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup, index int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello = ", index)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(&group, i)
	}

	group.Wait()
	fmt.Println("Selesai bro")
}
