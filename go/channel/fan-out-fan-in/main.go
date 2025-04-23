package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate scraping a URL.
func scrapeURL(url string) string {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // Simulate variable response time
	return fmt.Sprintf("Content from %s", url)
}

// Worker function that processes URLs and sends results to the results channel.
func worker(id int, urls <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		fmt.Printf("Worker %d scraping %s\n", id, url)
		result := scrapeURL(url)
		results <- result
	}
}

func main() {
	const numWorkers = 3
	urls := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
		"http://example.com/page5",
	}

	// Channels for URLs to scrape and for results
	urlsChan := make(chan string, len(urls))
	resultsChan := make(chan string, len(urls))

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start workers (Fan-Out)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, urlsChan, resultsChan, &wg)
	}

	// Send URLs to the urlsChan (Fan-Out)
	for _, url := range urls {
		urlsChan <- url
	}
	close(urlsChan) // No more URLs to send

	// Wait for all workers to finish
	wg.Wait()
	close(resultsChan) // No more results to collect

	// Collect and print results (Fan-In)
	for result := range resultsChan {
		fmt.Println(result)
	}

	fmt.Println("All URLs scraped.")
}
