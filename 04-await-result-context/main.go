package main

import (
	"context"
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	log.Print("Await result from another go routine and be responsive to a context")

	// Construct a context that will cancel after 500ms
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancelFunc() // avoid resource leak

	// Construct a result channel
	results := make(chan int)

	go func(res chan int) {
		log.Print("Starting some longer running task returning multiple values")
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			res <- rand.IntN(10)
		}
		close(res) // Close channel to indicate to the main goroutine that no more results are coming. Channel must be closed to avoid leaking resources.
		log.Print("Task complete")
	}(results) // Channels can be passed by value and still its "data stream" is shared

	log.Print("Awaiting task completion")

	// Goroutine may send multiple results, but need to consider context also
loop:
	for {
		select {
		case r := <-results:
			log.Printf("Main goroutine received a result: %d", r)
		case <-cancelCtx.Done():
			log.Print("Cancellation request")
			break loop
		}
	}

	log.Print("Main go routine completed")
}
