package main

import (
	"log"
	"math/rand/v2"
	"slices"
	"sync"
	"time"
)

type SharedData struct {
	resultValues []int
	sync.Mutex   // embed Mutex to obtain functionality for locking and unlocking the criticalValue
}

func NewSharedValue() *SharedData {
	return &SharedData{
		resultValues: make([]int, 0),
	}
}

func (s *SharedData) Add(value int) {
	s.Lock()
	defer s.Unlock()
	s.appendValue(value)
}

func (s *SharedData) Get() []int {
	s.Lock()
	defer s.Unlock()
	return slices.Clone(s.resultValues) // it would be quite dangerous to return the slice as is if later more methods were added that would modify the slice or if the caller was modifying the slice.
}

func (s *SharedData) appendValue(value int) {
	// WARNING - if tried to call s.Lock() again, it would be a deadlock. Mutex is not smart enough to know that the lock is already held!
	s.resultValues = append(s.resultValues, value) // this may allocate new memory for the internal array holding the data
}

// You may want to run such code with race detection -race flag, i.e go run -race .
func main() {
	log.Print("Updating shared value from multiple goroutines")

	// Construct a shared object
	shared := SharedData{}

	go func(data *SharedData) {
		log.Print("Starting some longer running task returning multiple values")
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			v := rand.IntN(10)
			log.Printf("task adding value %d", v)
			shared.Add(v)
		}
		log.Print("Task complete")
	}(&shared)

	log.Print("Sleeping for 0.5s to get some task results")
	time.Sleep(500 * time.Millisecond)

	firstResults := shared.Get()

	for i, r := range firstResults {
		log.Printf("First result(%d): %d", i, r)
	}

	// delay more to get more results
	time.Sleep(200 * time.Millisecond)
	secondResults := shared.Get()

	for i, r := range secondResults {
		log.Printf("Second result(%d): %d", i, r)
	}

	log.Print("Main goroutine has completed")
}
