package main

import (
	. "fmt"
	"leetcode/topics/t766/result"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		r := result.IsToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})
		Printf("%+v \n", r)
	}

	duration := time.Since(start).Nanoseconds()
	Printf("duration: %d ns\n", duration)
}
