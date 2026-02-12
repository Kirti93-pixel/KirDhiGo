package arrays

import "fmt"

func rotateImage(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		rev(&matrix[i])
	}
	return
}

func rev(arr *[]int) {
	start, end := 0, len(*arr)-1
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}

func Run_Rotate_Image() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("1. For matrix", matrix)
	rotateImage(matrix)
	fmt.Println("output is::", matrix)
	fmt.Println()

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	fmt.Println("2. For matrix", matrix)
	rotateImage(matrix)
	fmt.Println("output is::", matrix)
}
