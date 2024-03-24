package main

import (
	"sync"
	"testing"
)

// make large number of parallel writes to trigger "concurrent map writes" errors
func TestConcurrentMapWrites(t *testing.T) {
	store := NewInMemoryPlayerStore()
	player := "Pepper"
	winCount := 200

	var wg sync.WaitGroup
	wg.Add(winCount)

	for i := 0; i < winCount; i++ {
		go func(p string) {
			store.RecordWin(p)
			wg.Done()
		}(player)
	}
	wg.Wait()

	got := store.GetPlayerScore(player)
	if got != winCount {
		t.Errorf("did not get correct status, got %d, want %d", got, winCount)
	}
}
