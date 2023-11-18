/* HEAPSORT created using an array structure.
 * To see a btree based heap, check binary tree folder
 */

package main

import "fmt"

func rearrange(heap []int, root int, last int) {
	ok := false
	for !ok {
		leftChild := 2*root + 1
		if leftChild >= last {
			ok = true
		} else {
			rightChild := 2*root + 2
			iMax := root

			if rightChild < last && heap[rightChild] > heap[leftChild] {
				iMax = rightChild
			} else {
				iMax = leftChild
			}

			if heap[iMax] > heap[root] {
				heap[iMax], heap[root] = heap[root], heap[iMax]
				root = iMax
			} else {
				ok = true
			}
		}
	}
}

func buildHeap(values []int) []int {
	for i := len(values)/2 - 1; i >= 0; i-- {
		rearrange(values, i, len(values))
	}
	return values
}

func main() {
	values := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	heap := buildHeap(values)
	fmt.Println(heap)
}
