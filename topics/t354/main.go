package main

import (
	"fmt"
	"leetcode/topics/t354/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.MaxEnvelopes([][]int{{10, 4}, {13, 18}, {1, 5}, {13, 15}, {3, 12}, {12, 11}, {17, 15}, {7, 1}, {17, 18}, {7, 19}, {2, 5}, {8, 9}, {18, 10}, {7, 6}, {17, 7}})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
