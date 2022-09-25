package main

import (
	"fmt"
	"sort"
)

func findLongestSubsequence(arr []int32) int32 {
	// Write your code here

	// + sorted slices
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	s := sumDiff(arr)
	if isEven(s) {
		return int32(len(arr))
	}

	// condition for array probabilities
	var n int = len(arr) - 1
	for n > 0 {

		for i := 0; i <= n; i++ {
			arri := make([]int32, len(arr))
			// + copy slices
			copy(arri, arr)
			// + removed element in slices
			arrtemp := append(arri[:i], arri[i+1:]...)

			// + sum adjecent defferent variables
			s := sumDiff(arrtemp)

			// checked
			if isEven(s) {
				return int32(n)
			}

		}
		arr = arr[:n]
		n--
	}
	return 0
}

func sumDiff(arr []int32) (sum int32) {
	for i := 0; i < len(arr)-1; i++ {
		sum += (arr[i+1] - arr[i])
	}
	return sum
}

func isEven(sum int32) bool {
	return sum%2 == 0
}

var Case0 = []int32{7, 5, 6, 2, 3, 2, 4} // ans. 6
var Case1 = []int32{1, 3, 5, 7}          // ans. 4

func main() {

	// fmt.Println(findLongestSubsequence(Case0))

	fmt.Println(findLongestSubsequence(Case1))

}
