package main

import (
	"context"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	log.Print("Producers and consumers")

	producerCount := 2
	consumerCount := 3

	// Construct a communication channel between producers and consumers
	results := make(chan int, consumerCount)

	ctx, cancelFunc := context.WithCancel(context.Background())

	producerWg := sync.WaitGroup{}
	producerWg.Add(producerCount)
	for i := 0; i < producerCount; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup, n int) {
			defer wg.Done()
			tickDuration := 200 + rand.IntN(100)
			log.Printf("Starting producer %d with duration %dms", n, tickDuration)
			triggers := time.Tick(time.Duration(tickDuration) * time.Millisecond)
		loop:
			for {
				select {
				case <-triggers:
					v := rand.IntN(100)
					results <- v
					log.Printf("Producer %d sent value %d", n, v)
				case <-ctx.Done():
					log.Printf("Producer %d completed", n)
					break loop
				}
			}
		}(ctx, &producerWg, i)
	}

	consumerWg := sync.WaitGroup{}
	consumerWg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			for r := range results {
				log.Printf("Consumer %d received value %d", n, r)
			}
			log.Printf("Consumer %d complete", n)
		}(i, &consumerWg)
	}

	log.Printf("Main goroutine waiting for 1s")
	time.Sleep(time.Second) // give producers and consumers 1s to work
	log.Printf("Main goroutine stopping producers")
	cancelFunc()      // cancel using context
	producerWg.Wait() // wait for producers to terminate
	log.Printf("Main goroutine closing results channel")
	close(results)
	consumerWg.Wait() // wait for consumers to complete

	log.Print("Main go routine completed")
}
