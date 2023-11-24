package main

func longestValidParentheses(s string) int {
	var pila []string
	count := 0
	maxCount := 0
	for len(s) > 0 {
		char := s[0]
		s = s[1:]
		if char == '(' {
			pila = append(pila, "()")
		} else if len(pila) == 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		} else {
			pila = pila[1:]
			count += 2
		}
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount
}

/*
func main() {
	fmt.Println(longestValidParentheses("(()"))
}
*/
