package main

import (
	"fmt"
	"leetcode/topics/t59/result"
	"time"
)

func main() {
	r := make([][]int, 0)
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.GenerateMatrix(5)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
