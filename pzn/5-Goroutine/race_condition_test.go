package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", counter)

}
