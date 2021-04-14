package main

import (
	"fmt"
	"leetcode/topics/t208/result"
	"time"
)

func main() {
	start := time.Now()

	trie := &result.Trie{}
	trie.Insert("apple")
	fmt.Printf("%v\n", trie.Search("apple"))
	fmt.Printf("%v\n", trie.Search("app"))
	fmt.Printf("%v\n", trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Printf("%v\n", trie.Search("app"))

	duration := time.Since(start).Nanoseconds()
	fmt.Printf("duration: %d ns\n"+"result:%+v\n", duration, "done")
}
