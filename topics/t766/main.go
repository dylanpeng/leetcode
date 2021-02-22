package main

import (
	. "fmt"
	"leetcode/topics/t766/result"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		r := result.IsToeplitzMatrix([][]int{{83},{64},{57}})
		Printf("%+v \n", r)
	}

	duration := time.Since(start).Nanoseconds()
	Printf("duration: %d ns\n", duration)
}
