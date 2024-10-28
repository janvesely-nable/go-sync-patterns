package main

import (
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	log.Print("Await result from another go routine")

	// Construct a result channel
	results := make(chan int)

	go func(res chan int) {
		log.Print("Starting some longer running task")
		time.Sleep(time.Second)
		res <- rand.IntN(10)
		close(res) // Close channel to indicate to the main goroutine that no more results are coming. Channel must be closed to avoid leaking resources.
		log.Print("Task complete")
	}(results) // Channels can be passed by value and still its "data stream" is shared

	log.Print("Awaiting task completion")
	// Goroutine may send multiple results - so just loop through all results till the end of the channel.
	for r := range results {
		log.Printf("Main goroutine received result: %d", r)
	}

	log.Print("Main go routine completed")
}
