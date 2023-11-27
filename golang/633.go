package main

import (
	"math"
)

func judgeSquareSum(c int) bool {
	root := 0.
	end := int(math.Sqrt(float64(c)))
	for root != float64(end) {
		result := int(math.Pow(root, 2)) + int(math.Pow(float64(end), 2))
		if result == c {
			return true
		} else {
			if result > c {
				end--
			} else {
				root++
			}
		}
	}
	result := int(math.Pow(root, 2)) + int(math.Pow(float64(end), 2))
	if result == c {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(judgeSquareSum(3))
// }
