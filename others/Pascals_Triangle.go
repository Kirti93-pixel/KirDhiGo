package others

import (
	"fmt"
)

func generate(numRows int) [][]int {
    mainL := make([][]int, 0)
    L0 := []int{1}
    mainL = append(mainL, L0)
    if numRows == 1 {
        return mainL
    }
    L1 := []int{1,1}
    mainL = append(mainL, L1)
    if numRows == 2 {
        return mainL
    }
    if numRows >=3 {
        for i:=2; i<numRows;i++ {
            Li := make([]int, 0)
            Li = append(Li, 1)
            for j:=1;j<i;j++ {
                Li = append(Li, mainL[i-1][j-1]+mainL[i-1][j])
            }
            Li = append(Li, 1)
            mainL = append(mainL, Li)
        }
    }
    return mainL  
}

func Run_Pascals_Triangle() {
	fmt.Println("Make Pascals triangle for NumRows: 5", generate(5))
	fmt.Println("Make Pascals triangle for NumRows: 8", generate(8))
}

