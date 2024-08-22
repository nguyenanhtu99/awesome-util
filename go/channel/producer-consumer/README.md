# Producer/Consumer
## Problem
Imagine a factory where multiple machines (producers) generate parts (tasks), and a single worker (consumer) assembles these parts. We want to coordinate the machines and the worker so that the worker assembles parts as they are produced.
## Solution
1. Producer Function:
- The producer function generates a fixed number of tasks (random integers) and sends them to the jobs channel.
- Each producer simulates some work by sleeping for a random duration between tasks.
2. Consumer Function:
- The consumer function receives tasks from the jobs channel and processes them.
- Processing is simulated by sleeping for one second per task.
- After all tasks are processed (when the jobs channel is closed), the consumer sends a signal on the done channel to indicate it’s finished.
3. Main Function:
- We create a jobs channel to hold tasks and a done channel to signal when the consumer is finished.
- We start multiple producers, each running in its own goroutine.
- We start the consumer in its own goroutine.
- After all producers are done, we close the jobs channel, so the consumer knows when there are no more tasks.
- Finally, we wait for the consumer to signal that all tasks have been processed.

