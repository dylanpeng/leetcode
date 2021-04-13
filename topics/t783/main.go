package main

import (
	"fmt"
	"leetcode/topics/t783/result"
	"time"
)

func main() {
	r := 0
	start := time.Now()

	root := &result.TreeNode{
		Val: 90,
		Left: &result.TreeNode{
			Val: 69,
			Left: &result.TreeNode{
				Val: 49,
				Right: &result.TreeNode{
					Val: 52,
				},
			},
			Right: &result.TreeNode{
				Val: 89,
			},
		},
	}

	for i := 0; i < 1; i++ {
		r = result.MinDiffInBST(root)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
}
