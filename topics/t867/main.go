package main

import (
	"fmt"
	"leetcode/topics/t867/result"
	"time"
)

func main() {
	var r [][]int
	start := time.Now()

	for i := 0; i < 10; i++ {
		r = result.Transpose2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
