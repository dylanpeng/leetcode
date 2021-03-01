package main

import (
	"fmt"
	"leetcode/topics/t303/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	for i := 0; i < 10; i++ {
		obj := result.Constructor([]int{1, 2, 3, 4})
		r = obj.SumRange(1, 3)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
