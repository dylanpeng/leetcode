package main

import (
	"fmt"
	"leetcode/topics/t832/result"
	"time"
)

func main() {
	var r [][]int
	start := time.Now()

	for i := 0; i < 10; i++ {
		r = result.FlipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
