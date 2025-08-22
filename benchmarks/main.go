package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10_000_000

	// if-else
	start := time.Now()
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			// ...
		} else {
			// ...
		}
	}
	time1 := time.Since(start).Seconds()

	// separate if
	start = time.Now()
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			// ...
		}
		if i%2 != 0 {
			// ...
		}
	}
	time2 := time.Since(start).Seconds()

	fmt.Printf("if-else: %.6fs\n", time1)
	fmt.Printf("separate if: %.6fs\n", time2)
	fmt.Printf("Difference: %.6fs\n", time1-time2)
	fmt.Printf("%% Faster: %.2f%%\n", ((time1-time2)/time2)*100)
}
