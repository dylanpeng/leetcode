package main

import (
	"fmt"
	"leetcode/topics/t1004/result"
)

func main() {
	for i:= 0; i < 10; i++ {
		r := result.LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
		fmt.Printf("%d \n", r)
	}
}
