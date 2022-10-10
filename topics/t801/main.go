package main

import (
	"fmt"
	result2 "leetcode/topics/t801/result"
)

func main() {
	nums1 := []int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}
	nums2 := []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}
	result := result2.MinSwap(nums1, nums2)
	fmt.Printf("%+v, %+v, %d\n", nums1, nums2, result)
}
