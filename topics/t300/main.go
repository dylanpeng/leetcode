package main

import (
	"fmt"
	"leetcode/topics/t300/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.LengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
