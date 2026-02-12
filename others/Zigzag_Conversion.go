package others

import (
	"fmt"
	"strings"
)

func zigzagConversion(s string, numRows int) string {
	zigArr := make([][]string, numRows)
	lenS := len(s)
	j := -1
	direction := "forward"
    if numRows == 1 {
        return s
    }
	for i := 0; i < lenS; i++ {
		if direction == "forward" {
			j++
			if j == numRows - 1 {
				direction = "backward"
			} 
            zigArr[j] = append(zigArr[j], string(s[i]))
		} else if direction == "backward" {
			j--
			if j == 0 {
				direction = "forward"
			} 
            zigArr[j] = append(zigArr[j], string(s[i]))
		}
	}

	res := ""
	for k := 0; k < numRows; k++ {
		res = res + strings.Join(zigArr[k], "")
	}

	return res
}

func Run_Zigzag_Conversion() {
	fmt.Println("for s=\"PAYPALISHIRING\" and numRows=3, output is::", zigzagConversion("PAYPALISHIRING", 3))
	fmt.Println("for s=\"PAYPALISHIRING\" and numRows=4, output is::", zigzagConversion("PAYPALISHIRING", 4))
}