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
	sync.RWMutex // embed Mutex to obtain functionality for locking and unlocking the criticalValue
}

func NewSharedValue() *SharedData {
	return &SharedData{
		resultValues: make([]int, 0),
	}
}

func (s *SharedData) Add(value int) {
	s.Lock() // lock for writing
	defer s.Unlock()
	s.appendValue(value)
}

func (s *SharedData) Get() []int {
	s.RLock() // lock for reading
	defer s.RUnlock()
	return slices.Clone(s.resultValues) // it would be quite dangerous to return the slice as is if later more methods were added that would modify the slice or if the caller was modifying the slice.
}

func (s *SharedData) appendValue(value int) {
	// WARNING - if tried to call s.Lock() or s.RLock() again, it would be a deadlock. Mutex is not smart enough to know that the lock is already held!
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
			time.Sleep(10 * time.Millisecond)
			v := rand.IntN(10)
			log.Printf("task adding value %d", v)
			shared.Add(v)
		}
		log.Print("Task complete")
	}(&shared)

	// spin-up multiple readers
	wg := sync.WaitGroup{}
	readerCount := 5
	wg.Add(readerCount)
	for i := 0; i < readerCount; i++ {
		go func(n int, data *SharedData, wg *sync.WaitGroup) {
			defer wg.Done()
			dur := rand.IntN(50) + 10
			log.Printf("Reader %d waiting for %dms", i, dur)
			time.Sleep(time.Duration(dur * int(time.Millisecond)))
			values := shared.Get()
			log.Printf("Reader %d has read %v values", n, values)
		}(i, &shared, &wg)
	}

	wg.Wait()

	log.Print("Main goroutine has completed")
}
