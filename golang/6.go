package main 

func convert(s string, numRows int) (res string) {
	if numRows == 1 {
		return s
	}
	var matrix [][]string
	for i := 0; i < numRows; i++ {
		row := make([]string, len(s))
		matrix = append(matrix, row)
	}

	row, col, direction := 0, 0, 1
	for _, char := range s {
		matrix[row][col] = string(char)
		col++
		if row == numRows-1 {
			direction = -1
		} else if row == 0 {
			direction = 1
		}
		row += direction
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != "" {
				res += matrix[i][j]
			}
		}
	}
	return res
}
