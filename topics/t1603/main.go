package main

import (
	"fmt"
	"leetcode/topics/t1603/result"
	"time"
)

func main() {
	r := result.Constructor(1, 2, 3)
	start := time.Now()

	for i := 0; i < 1; i++ {
		_ = r.AddCar(1)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
