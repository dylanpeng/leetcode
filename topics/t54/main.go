package main

import (
	"fmt"
	"leetcode/topics/t54/result"
	"time"
)

func main() {
	r := make([]int, 0)
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
