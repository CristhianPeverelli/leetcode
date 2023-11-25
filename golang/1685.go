package main

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	cumulativeSum := make([]int, n)
	cumulativeSum[0] = nums[0]
	for i := 1; i < n; i++ {
		cumulativeSum[i] = cumulativeSum[i-1] + nums[i]
	}

	for i := 0; i < n; i++ {
		leftSum := cumulativeSum[i] - nums[i]
		rightSum := cumulativeSum[n-1] - cumulativeSum[i]

		result[i] = (nums[i]*i - leftSum) + (rightSum - nums[i]*(n-i-1))
	}

	return result
}

// func main() {
// 	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))
// 	fmt.Println(getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10}))
// 	fmt.Println(getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10, 3, 9}))
// }
