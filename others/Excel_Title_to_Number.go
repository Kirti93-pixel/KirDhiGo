package others

import "fmt"

func titleToNumber(columnTitle string) int {
    res := 0
    for _, s := range columnTitle {
        res = res * 26 + (int(rune(s)) - int('A')) + 1
    }
    return res
}

func Run_Excel_Title_to_Number() {
	fmt.Println("Title BGDH is", titleToNumber("BGDH"))
	fmt.Println("Title ZY is", titleToNumber("ZY"))
	fmt.Println("Title U is", titleToNumber("U"))
}