package others

import "fmt"

func convertToTitle(columnNumber int) string {
    res := ""
    var A rune = 'A'
    for columnNumber>0{
        columnNumber-=1
        res = string(rune(int(A)+columnNumber%26)) + res
        columnNumber/=26
    }
    return res
}

func Run_Covert_excel_column_to_Alphabet_Name() {
	fmt.Println("Column 701 is", convertToTitle(701))
	fmt.Println("Column 85247 is", convertToTitle(85247))
	fmt.Println("Column 28 is", convertToTitle(28))
	fmt.Println("Column 7 is", convertToTitle(7))
}