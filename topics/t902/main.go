package main

import (
	"fmt"
	"leetcode/topics/t902/result"
)

func main() {
	nList, n := result.LenInt(1000000000)

	fmt.Printf("%+v, %d \n", nList, n)

	list, l := result.Sort902([]string{"1", "4", "9"})

	fmt.Printf("%+v, %d \n", list, l)

	count := result.AtMostNGivenDigitSet([]string{"3", "4", "8"}, 4)

	fmt.Printf("count:%d \n", count)

	return
}
