package benchmark

import (
	"fmt"
	"time"

	"github.com/knork-fork/go-benchmarker/pkg/request"
)

func StartBenchmark(threadCount int, requestsPerThread int, url string) {
	for i := 0; i < threadCount; i++ {
		go StartBenchmarkOnThread(i, requestsPerThread, url)
	}
}

func StartBenchmarkOnThread(thread int, requestCount int, url string) {
	var failedRequests int = 0
	var startTime time.Time = time.Now()

	for i := 0; i < requestCount; i++ {
		if !request.Send(url) {
			failedRequests++
		}
	}

	var durationTime = time.Since(startTime)

	fmt.Printf("Finished %v requests on thread #%v (%v failed) in %vms\n", requestCount, thread, failedRequests, durationTime.Milliseconds())
}
