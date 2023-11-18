package main

func quickSortAlt(values []int) []int {
	if len(values) > 1 {
		pivotIndex := len(values) - 1
		pivot := values[pivotIndex]
		a := []int{}
		b := []int{}
		for _, val := range values[:pivotIndex] {
			if val <= pivot {
				a = append(a, val)
			} else {
				b = append(b, val)
			}
		}
		a = quickSort(a)
		b = quickSort(b)
		values = append(append(a, pivot), b...)
	}
	return values
}

func quickSortRec(values []int, i int, j int) []int {
	return values
}

func quickSort(values []int) []int {
	values = quickSortRec(values, 0, len(values))
	return values
}
