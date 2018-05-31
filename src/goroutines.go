package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib_worker(n int) {
	defer wg.Done()
	fib(n)
}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

// START OMIT
func main() {
	defer trackTime(time.Now(), "MAIN") // HL
	MAX := 64
	FIB := 35
	for i := 0; i < MAX; i++ {
		wg.Add(1)
		fib_worker(FIB) // HL
	}
	wg.Wait()
	fmt.Println("Total Times FIB Executed:", MAX)
}

// END OMIT
