# Fan Out/Fan In
## Problem
Imagine you need to scrape data from multiple websites concurrently. You can fan out the work by assigning different URLs to different goroutines, and then fan in the results by collecting data from all the goroutines.
## Solution
1. Worker Function:
- The worker function receives URLs from the urls channel, scrapes them using the scrapeURL function, and sends the result to the results channel.
- The sync.WaitGroup is used to track when each worker has finished its tasks.
2.	Main Function:
- We create a list of URLs to scrape.
- We create two channels: urlsChan to distribute URLs to workers and resultsChan to collect the scraped content.
- We start a fixed number of workers, each in its own goroutine, to scrape URLs (Fan-Out).
- URLs are sent to the urlsChan, and the channel is closed after all URLs are sent.
- We wait for all workers to finish using wg.Wait() and then close the resultsChan.
- Finally, we read and print the results from the resultsChan (Fan-In).
