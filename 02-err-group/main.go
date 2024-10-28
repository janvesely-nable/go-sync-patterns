package main

import (
	"errors"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.Print("Await completion of another go routine")

	// Construct a new errgroup
	eg := errgroup.Group{}

	eg.Go(func() error {
		log.Print("Starting some longer running task")
		time.Sleep(2 * time.Second)
		log.Print("Task completed with error")
		return errors.New("goroutine badly failed")
	})

	log.Print("Awaiting task completion")

	err := eg.Wait() // wait for all - note that it returns the first error should there be more than 1
	if err != nil {
		log.Printf("Goroutine failed with: %v", err)
	}

	log.Print("Main go routine completed its wait on task")
}
