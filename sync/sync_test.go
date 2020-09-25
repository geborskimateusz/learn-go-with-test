package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing counter leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("run concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()
		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("Expected %d, got %d", want, got.Value())
	}
}
