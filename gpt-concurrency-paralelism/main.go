package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyComputation(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("Task %d started\n", id)

	sum := 0
	for i := 0; i < 1e9; i++ {
		sum += i
	}

	fmt.Printf("Task %d completed in %v\n", id, time.Since(start))
}

func main() {
	// 1. Setting it to one ensures that we're actually using Concurrency - N tasks running while competing for
	// a single CPU Core. So the manager will schedule time for each task to be processed.
	// 2. Setting it to N > 1 allows the tasks to be processed in parallel, parallelism, which is not the same as concurrency.
	runtime.GOMAXPROCS(12)

	var wg sync.WaitGroup
	wg.Add(2) // Number of tasks added into the WaitGroup

	start := time.Now()

	go heavyComputation(1, &wg)
	go heavyComputation(2, &wg)

	wg.Wait() // Wait for all tasks to complete

	fmt.Printf("Total time taken: %v\n", time.Since(start))
}
