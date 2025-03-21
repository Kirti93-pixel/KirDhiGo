package arrays

import (
	"fmt"
	"strconv"
)

func setMatrixZeroes(matrix [][]int) {
	setZeroes1(matrix)
}

func setZeroes1(matrix [][]int) [][]int {
	r, c, fr, fc := len(matrix), len(matrix[0]), false, false
	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			fc = true
		}
	}
	for j := 0; j < c; j++ {
		if matrix[0][j] == 0 {
			fr = true
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < r; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < c; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < c; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < r; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for j := 0; j < c; j++ {
			matrix[0][j] = 0
		}
	}

	if fc {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}

	return matrix
}

func setZeroes2(matrix [][]int) [][]int {
	chile0PosStore := map[string]struct{}{}

	rowLen, colLen := len(matrix), len(matrix[0])

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			iS, jS := strconv.Itoa(i), strconv.Itoa(j)
			if _, exists := chile0PosStore[iS+":"+jS]; exists {
				continue
			}
			if matrix[i][j] != 0 {
				continue
			}
			for k := 0; k < rowLen; k++ {
				if matrix[k][j] != 0 {
					matrix[k][j] = 0
					kS, jS := strconv.Itoa(k), strconv.Itoa(j)
					chile0PosStore[kS+":"+jS] = struct{}{}
				}
			}
			for k := 0; k < colLen; k++ {
				if matrix[i][k] != 0 {
					matrix[i][k] = 0
					iS, kS := strconv.Itoa(i), strconv.Itoa(k)
					chile0PosStore[iS+":"+kS] = struct{}{}
				}
			}
		}
	}
	return matrix
}

func Run_Set_Matrix_Zeroes() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println("1. For matrix", matrix, " output is::")
	setMatrixZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	fmt.Println("2. For matrix", matrix, " output is::")
	setMatrixZeroes(matrix)
	fmt.Println(matrix)
}
