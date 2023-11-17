package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(a1 []int, a2 []int) []int {
	supp := []int{}
	i1, i2 := 0, 0
	for i1 < len(a1) && i2 < len(a2) {
		if a1[i1] < a2[i2] {
			supp = append(supp, a1[i1])
			i1++
		} else {
			supp = append(supp, a2[i2])
			i2++
		}
	}
	supp = append(supp, a1[i1:]...)
	supp = append(supp, a2[i2:]...)
	return supp
}
