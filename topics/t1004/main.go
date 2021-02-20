package main

import (
	"fmt"
	"leetcode/topics/t1004/result"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		r := result.LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
		fmt.Printf("%d \n", r)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n", duration)
}
