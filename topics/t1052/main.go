package main

import (
	"fmt"
	"leetcode/topics/t1052/result"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		r := result.MaxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)
		fmt.Printf("%d \n", r)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n", duration)
}
