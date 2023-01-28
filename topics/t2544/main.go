package main

import (
	"fmt"
	"leetcode/topics/t2544/result"
)

func main() {
	input := 88699612
	count := result.AlternateDigitSum(input)

	fmt.Printf("input: %d | out:%d\n", input, count)
}
