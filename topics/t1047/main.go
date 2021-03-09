package main

import (
	"fmt"
	"leetcode/topics/t1047/result"
	"time"
)

func main() {
	r := ""
	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.RemoveDuplicates("ibfjcaffccadidiaidchakchchcahabhibdcejkdkfbaeeaecdjhajbkfebebfea")
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
