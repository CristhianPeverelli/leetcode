package main

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		selectedIndex := 0
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[selectedIndex] {
				selectedIndex = j
			}
		}
		arr[selectedIndex], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[selectedIndex]
	}
}
