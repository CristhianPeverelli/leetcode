package main

func removeDuplicateLetters(s string) string {
	lastIndex := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]] = i
	}

	stack := make([]byte, 0)
	visited := make(map[byte]bool)

	for i, char := range s {
		n := byte(char)
		if visited[n] {
			continue
		}
		for len(stack) > 0 && n < stack[len(stack)-1] && i < lastIndex[stack[len(stack)-1]] {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, n)
		visited[n] = true

	}

	return string(stack)
}

// func main() {
// 	fmt.Println(removeDuplicateLetters("jdjcnwjqj"))
// }
