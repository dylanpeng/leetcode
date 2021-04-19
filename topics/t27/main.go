package main

import (
	"fmt"
	"leetcode/topics/t27/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.RemoveElement([]int{3,3,3}, 3)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}