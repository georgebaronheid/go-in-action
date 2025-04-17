package main

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go heavyComputation(1, &wg)
		go heavyComputation(2, &wg)
		wg.Wait()
	}
}

func BenchmarkParallel(b *testing.B) {
	runtime.GOMAXPROCS(6)
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go heavyComputation(1, &wg)
		go heavyComputation(2, &wg)
		wg.Wait()
	}
}
