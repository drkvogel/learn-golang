package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func consume(nGoroutine int) {
	fmt.Println("Starting Goroutine", nGoroutine)
	n := time.Now().UnixNano()
	for {
		n++
	}
}
func main() {
	log.Println("NumCpu", runtime.NumCPU())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go consume(i)
	}
	wg.Wait()
}
