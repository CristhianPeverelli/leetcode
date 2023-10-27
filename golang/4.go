package main

import (
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) (median float64) {
	merged := append(nums1, nums2...)
	slices.Sort(merged)
	if len(merged)%2 == 0 {
		median = float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2
	} else {
		median = float64(merged[len(merged)/2])
	}
	return
}
