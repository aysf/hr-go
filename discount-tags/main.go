package main

import "fmt"

func GetMaxEvenSum(val []int32) int64 {

	var sum int32
	for _, v := range val {
		if v >= 0 {
			sum += v
		}
	}
	if isEven(sum) {
		return int64(sum)
	}

	return int64(sum - abs(getMinOdd(val)))

}

func getMinOdd(val []int32) int32 {

	var min int32
	if val[0] != 0 {
		min = val[0]
	} else {
		min = 3
	}
	for _, v := range val {
		if abs(v) < abs(min) && !isEven(abs(v)) {
			min = v
		}
	}
	return min
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

var Case0 = []int32{2, 3, 6, -5, 10, 1, 1} // ans. 22
var Case1 = []int32{-1, -2, -3, 8, 7}      // ans. 14
var Case2 = []int32{2, 3, 6, -5, 10, 1, 1} // ans. 22
var Case3 = []int32{6, 3, 4, -1, 9, 17}    // ans. 38
var Case4 = []int32{91}                    // ans. 38

func main() {
	fmt.Println(GetMaxEvenSum(Case0))
	fmt.Println(GetMaxEvenSum(Case1))
	fmt.Println(GetMaxEvenSum(Case2))
	fmt.Println(GetMaxEvenSum(Case3))
	fmt.Println(GetMaxEvenSum(Case4))
}
