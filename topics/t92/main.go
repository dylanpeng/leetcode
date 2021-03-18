package main

import (
	"fmt"
	"leetcode/topics/t92/result"
	"time"
)

func main() {
	r := &result.ListNode{}
	head := &result.ListNode{
		Val:  1,
		Next: nil,
	}

	tail := head

	for i := 2; i <= 5; i++ {
		node := &result.ListNode{
			Val:  i,
			Next: nil,
		}
		tail.Next = node
		tail = node
	}

	tail = head
	for ; tail != nil; tail = tail.Next {
		fmt.Printf("%d ", tail.Val)
	}
	fmt.Println()

	start := time.Now()

	for i := 0; i < 1; i++ {
		r = result.ReverseBetween(head, 4, 5)
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, r)
	tail = r
	for ; tail != nil; tail = tail.Next {
		fmt.Printf("%d ", tail.Val)
	}
	fmt.Println()
}
