package main

func longestPalindrome(s string) (longest string) {
	for i, _ := range s {
		for j, _ := range s {
			if i <= j && (len(longest) <= len(s[i:j+1])) && isPalindrome(s[i:j+1]) {
				longest = s[i : j+1]
			}

		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
