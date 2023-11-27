package main

func nextGreatestLetter(letters []byte, target byte) byte {
	var greatest byte = '{'
	for _, char := range letters {
		if char > target && char < greatest {
			greatest = char
		}
	}
	if greatest == '{' {
		return letters[0]
	}
	return greatest
}

// func main() {
// 	a := []byte{'x', 'x', 'y'}
// 	fmt.Println(string(nextGreatestLetter(a, 'z')))
// }
