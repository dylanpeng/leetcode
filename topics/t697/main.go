package main

import (
	"fmt"
	"leetcode/topics/t697/result"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		r := result.FindShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})
		fmt.Printf("%d \n", r)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n", duration)
}
