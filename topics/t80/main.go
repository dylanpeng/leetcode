package main

import (
	"fmt"
	"leetcode/topics/t80/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()
	nums := []int{0,0,1,1,1,1,2,3,3}

	for i := 0; i < 1; i++ {
		r = result.RemoveDuplicates(nums)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
