package main

import (
	"fmt"
	"slices"
)

func maxCoins(piles []int) int {
	count := 0
	slices.Sort(piles)
	fmt.Println(piles)
	index := len(piles) - 2
	for index > (len(piles)/3 - 1) {
		count += piles[index]
		index -= 2
	}
	return count
}

/* ALTERNATIVE WITH REMOVAL:
package main

import (
	"fmt"
	"slices"
)

func maxCoins(piles []int) int {
	count := 0
	slices.Sort(piles)
	for len(piles) > 0 {
		piles = remove(piles, len(piles)-1)
		count += piles[len(piles)-1]
		piles = remove(piles, len(piles)-1)
		piles = remove(piles, 0)
	}
	return count
}

func remove(values []int, index int) []int {
	if index == len(values)-1 {
		values = values[:index]
	} else {
		values = append(values[:index], values[index+1:]...)
	}
	return values
}

func main() {
	piles := []int{2, 4, 1, 2, 7, 8}
	fmt.Println(maxCoins(piles))
}

*/
