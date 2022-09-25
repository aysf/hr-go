package main

import (
	"fmt"
	"sort"
)

func findNumOfPairs(a []int32, b []int32) int32 {
	// Write your code here

	// sorted slice a
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	// sorted slice b
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	n := len(a)
	if len(a) > len(b) {
		n = len(b)
	}
	count := 0
	for i := 0; i < n; i++ {
		for idxB, valB := range b {
			if valB < a[i] {
				count++
				b = b[idxB+1:]
				fmt.Println("iterasi:", i, ":", a[i], valB)
				fmt.Println(b)
				break
			}
		}
	}

	return int32(count)
}

var (
	// Case1a = []int32{2, 3, 1, 5, 4}
	// 5, 4, 3, 2, 1
	// 6, 6, 1, 1, 1
	// Case1b = []int32{1, 6, 1, 6, 1} // ans. 3

	Case2a = []int32{2, 1, 0, 0}
	Case2b = []int32{3, 1, 0} // ans. 2
)

func main() {
	// fmt.Println(findNumOfPairs(Case1a, Case1b))

	fmt.Println(findNumOfPairs(Case2a, Case2b))
}
