package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort() {
	reader := bufio.NewReader(os.Stdin)
	array := []int{}
	for {
		fmt.Print("Enter a number (or 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number or 'exit' to quit.")
		} else {
			inserted := false
			for i, n := range array {
				if number < n {
					secondPart := append([]int{}, array[i:]...)
					array = append(array[:i], number)
					array = append(array, secondPart...)
					inserted = true
					break
				}
			}
			if !inserted {
				array = append(array, number)
			}
			fmt.Printf("Inserted %d to the list\n", number)
			fmt.Println("List now is: ", array)
		}
	}
}
