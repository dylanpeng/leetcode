package main

import (
	"fmt"
	"leetcode/topics/t456/result"
	"time"
)

func main() {
	r := false
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.Find132pattern2([]int{-2, 1, 2, -2, 1, 2})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
