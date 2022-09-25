package main

import (
	"fmt"
	"math"
	"sort"
)

func getMaximumScore(arr []int32, k int32) int64 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var sum int32
	last := len(arr) - 1
	var third float64 = 3
	for i := 1; i <= int(k); i++ {
		sum += arr[last]

		n := math.Ceil(float64(arr[last]) / third)

		// + removed last element in slices
		arr = arr[:last]

		// + insert element in slices
		arr = inserter(int32(n), arr)

	}
	return int64(sum)
}

func inserter(val int32, arr []int32) []int32 {
	if arr[len(arr)-1] < val {
		return append(arr, val)
	}
	for i, el := range arr {
		if val <= el {
			arr = insert(arr, val, i)
			break
		}
	}
	return arr
}

func insert(a []int32, c int32, i int) []int32 {
	return append(a[:i], append([]int32{c}, a[i:]...)...)
}

var Case3arr = []int32{20, 4, 3, 1, 9} // ans. 40
var Case3k = 4
var Case0arr = []int32{4, 5, 18, 1} // ans. 29
var Case0k = 3
var Case1arr = []int32{1, 1, 1} // ans. 2
var Case1k = 2

func main() {

	var t int64

	t = getMaximumScore(Case0arr, int32(Case0k))
	fmt.Println(t)

	t = getMaximumScore(Case1arr, int32(Case1k))
	fmt.Println(t)

	t = getMaximumScore(Case3arr, int32(Case3k))
	fmt.Println(t)
}
