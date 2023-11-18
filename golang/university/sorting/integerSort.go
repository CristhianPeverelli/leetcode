package main

import "fmt"

func integerSort(values []int) []int {
	if len(values) == 0 {
		return nil
	}
	max := values[0]
	for _, num := range values {
		if max < num {
			max = num
		}
	}
	sorted := make([]int, max+1)
	for _, num := range values {
		sorted[num]++
	}
	return sorted
}

func printIS(sorted []int) {
	for j, num := range sorted {
		if num > 0 {
			for i := 0; i < num; i++ {
				fmt.Print(j, ",")
			}
		}
	}
}

/*
func main() {
	arr := []int{5, 2, 3, 7, 36, 12, 45, 63, 12}
	fmt.Println(arr)
	a := integerSort(arr)
	printIS(a)
} */
