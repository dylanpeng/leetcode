package main

import (
	"fmt"
	"leetcode/topics/t132/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.MinCut("cabababcbc")
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
