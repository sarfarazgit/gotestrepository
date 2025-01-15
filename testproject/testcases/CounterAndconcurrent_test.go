package main

import (
	"sync"
	"testing"
)

type Counter struct {
	mutex sync.Mutex
	Value int
}

func (count *Counter) Increament() {
	count.mutex.Lock()
	defer count.mutex.Unlock()
	count.Value++
}
func (count *Counter) TotalCounts() int {
	return count.Value
}

func TestCounter(t *testing.T) {
	t.Run("Increament the counter 3 times", func(t *testing.T) {
		counter := Counter{Value: 0}
		for i := 0; i < 3; i++ {
			counter.Increament()
		}

		if counter.TotalCounts() != 3 {
			t.Errorf("got %d, want %d", counter.TotalCounts(), 3)

		}
	})
}

func TestSafeConcurrentUsingMutex(t *testing.T) {

	t.Run("Safe concurrent test", func(t *testing.T) {
		expectedcount := 1000
		var waitgroup sync.WaitGroup
		counter := Counter{Value: 0}

		waitgroup.Add(expectedcount)

		for i := 0; i < expectedcount; i++ {
			go func() {
				counter.Increament()
				waitgroup.Done()
			}()
		}
		waitgroup.Wait()

		if counter.TotalCounts() != expectedcount {
			t.Errorf("The result %d, want %d", counter.TotalCounts(), expectedcount)
		}
	})

}
