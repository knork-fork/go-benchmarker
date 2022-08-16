package main

import (
	"time"

	test "github.com/knork-fork/go-benchmarker/pkg/testsrc"
)

func main() {
	test.TestMessage("Hello world!")
	// Sleep for 2 seconds
	time.Sleep(time.Second * 2)
}
