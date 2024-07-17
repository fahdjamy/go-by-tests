package concurrency

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("counter increments correctly", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("expected value %d, got %d", 3, counter.Value())
		}
	})

	t.Run("counter must be safe when called concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)
		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		if counter.Value() != expectedCount {
			t.Errorf("expected value %d, got %d", expectedCount, counter.Value())
		}
	})
}
