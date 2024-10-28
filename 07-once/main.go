package main

import (
	"log"
	"sync"
)

func OnlyRunMeOnce() {
	log.Printf("OnlyRunMeOnce running")
}

func main() {
	log.Print("Executing a function once")

	// Construct a wrapper for invoking OnlyRunMeOnce
	once := sync.Once{}

	once.Do(OnlyRunMeOnce) // first invocation will execute the function
	once.Do(OnlyRunMeOnce) // subsequent invocations will supress its execution

	log.Print("Main has completed")
}
