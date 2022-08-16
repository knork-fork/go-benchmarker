package main

import (
	"fmt"

	"github.com/knork-fork/go-benchmarker/pkg/benchmark"
)

func main() {
	// Run 1000 requests on each of 12 threads
	benchmark.StartBenchmark(12, 1000, "http://localhost:80/")

	// To-do: use sync.WaitGroup to wait for all threads to finish
	fmt.Scanln()
}
