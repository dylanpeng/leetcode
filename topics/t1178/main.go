package main

import (
	"fmt"
	"leetcode/topics/t1178/result"
	"time"
)

func main() {
	var r []int
	start := time.Now()

	for i := 0; i < 10; i++ {
		r = result.FindNumOfValidWords2([]string{"apple", "pleas", "please"}, []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"})
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
