package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Print("Await completion of another go routine")

	// Construct a new waitgroup
	wg := sync.WaitGroup{}

	wg.Add(1) // this waitgroup will have 1 task to wait for

	go func(wg *sync.WaitGroup) {
		defer wg.Done() // signal that one activity of the waitgroup is complete
		log.Print("Starting some longer running task")
		time.Sleep(2 * time.Second)
		log.Print("Task complete")
	}(&wg) // pass in the pointer to the waitgroup

	log.Print("Awaiting task completion")

	wg.Wait() // block until wg.Done() has been called by the other go routine

	log.Print("Main go routine completed its wait on task")
}
