package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	requestCount int64 = 0
	highestScore int64 = 100
	systemStatus int64 = 0
)

func worker(wg *sync.WaitGroup, score int64) {
	defer wg.Done()
	atomic.AddInt64(&requestCount, 1)

	current := atomic.LoadInt64(&highestScore)

	if score > current {
		atomic.CompareAndSwapInt64(&highestScore, current, score)
	}
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup

	atomic.StoreInt64(&systemStatus, 1)

	scores := []int64{120, 90, 20, 150, 200, 10, 2}

	for _, score := range scores {
		wg.Add(1)
		go worker(&wg, score)
	}

	wg.Wait()

	fmt.Println("System Status:", atomic.LoadInt64(&systemStatus))
	fmt.Println("Request Count:", atomic.LoadInt64(&requestCount))
	fmt.Println("Highest Score:", atomic.LoadInt64(&highestScore))
}
