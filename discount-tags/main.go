package main

import "fmt"

func getMaximumEvenSum(val []int32) int64 {
	// Write your code here

	// + sorted slices based on absolute value
	// slices (2, -4, -2, -1, 3) become (-1, -2, 2, 3, -4)
	sortedArray := groupSpc(val)

	var sum int32
	for _, se := range sortedArray {
		if se > 0 {
			sum += se
		}
	}

	if isEven(sum) {
		return int64(sum)
	}
	for _, m := range sortedArray {
		if !isEven(abs(m)) {
			return int64(sum - abs(sum))
		}
	}
	return 0
}

func groupSpc(val []int32) (res []int32) {
	for _, v := range val {

		res = append(res, v)
		if lr := len(res); lr > 1 {
			if abs(res[lr-2]) > abs(res[lr-1]) {
				res[lr-2], res[lr-1] = res[lr-1], res[lr-2]
			}
		}
	}
	return
}

func isEven(sum int32) bool {
	return sum%2 == 0
}

func abs(v int32) int32 {
	if v < 0 {
		return -1 * v
	}
	return v
}

var Case1 = []int32{5, -1, -2, -3, 8, 7}      // ans. 14
var Case2 = []int32{7, 2, 3, 6, -5, 10, 1, 1} // ans. 22

func main() {
	fmt.Println(getMaximumEvenSum(Case1))
	fmt.Println(getMaximumEvenSum(Case2))
}
