package booking

import (
	"sync"
	"sync/atomic"
	"testing"
	"github.com/google/uuid"
)

func TestConcurrentBooking_ExactlyOneWins(t *testing.T) {
	store := NewConcurrentStore()
	svc := NewService(store)

	const numGoroutines = 100_000

	var (
		successes atomic.Int64
		failures  atomic.Int64
		wg        sync.WaitGroup
	)

	wg.Add(numGoroutines)
	for i := range numGoroutines {
		go func(userNum int) {
			defer wg.Done()
			_, err := svc.Book(Booking{
				FlightID: "FL-NYC-LAX-001",
				SeatID:   "A1",
				UserID:   uuid.New().String(),
			})
			if err == nil {
				successes.Add(1)
			} else {
				failures.Add(1)
			}
		}(i)
	}

	wg.Wait()

	if successes.Load() != 1 {
		t.Errorf("expected exactly 1 success got %d", successes.Load())
	}
	if failures.Load() != numGoroutines-1 {
		t.Errorf("expected %d failures got %d", numGoroutines-1, failures.Load())
	}
}