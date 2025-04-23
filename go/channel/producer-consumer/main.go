package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function that generates tasks and sends them to the jobs channel.
func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		task := rand.Intn(100) // Generate a random task (an integer)
		fmt.Printf("Producer %d produced task %d\n", id, task)
		jobs <- task
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
	}
}

// Consumer function that processes tasks from the jobs channel.
func consumer(jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Consumer processing task %d\n", job)
		time.Sleep(time.Second) // Simulate processing time
	}
	done <- true // Signal that the consumer is done processing
}

func main() {
	const numProducers = 3

	// Channel for jobs and a channel to signal when consumer is done
	jobs := make(chan int, 10)
	done := make(chan bool)

	// WaitGroup to wait for all producers to finish
	var wg sync.WaitGroup

	// Start producers
	for p := 1; p <= numProducers; p++ {
		wg.Add(1)
		go producer(p, jobs, &wg)
	}

	// Start consumer
	go consumer(jobs, done)

	// Wait for all producers to finish
	wg.Wait()
	close(jobs) // Close the jobs channel to signal no more tasks

	// Wait for the consumer to finish
	<-done
	fmt.Println("All tasks processed.")
}
