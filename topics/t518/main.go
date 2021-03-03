package main

import (
	"fmt"
	"leetcode/topics/t518/result"
	"time"
)

func main() {
	r := make([]int, 0)
	start := time.Now()

	for i := 0; i < 10; i++ {
		r = result.CountBits2(5)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
