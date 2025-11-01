package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker represents a worker that processes jobs
type Worker struct {
	id int
}

// Process simulates doing work on a job, waits for a random delay, then sends the result
func (w Worker) Process(job int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()          // mark job as done when function exits
	delay := rand.Intn(1000) // random delay between 0–999 ms
	time.Sleep(time.Duration(delay) * time.Millisecond)
	results <- fmt.Sprintf("Worker %d finished job %d in %dms", w.id, job, delay)
}

func main() {
	numWorkers := 4
	numJobs := 10

	// channels for sending jobs and receiving results
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	// start a pool of worker goroutines
	for i := 1; i <= numWorkers; i++ {
		worker := Worker{id: i}
		go func(w Worker) {
			for job := range jobs { // continuously listen for new jobs
				wg.Add(1) // increment the WaitGroup for each job
				w.Process(job, &wg, results)
			}
		}(worker)
	}

	// send jobs into the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // close channel so workers stop when all jobs sent

	// wait until all jobs are processed
	wg.Wait()
	close(results) // close results channel after all done

	// print all results
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("All jobs completed")
}
