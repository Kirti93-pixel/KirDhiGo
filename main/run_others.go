package main

import (
	"fmt"

	"github.com/Kirti93-pixel/KirDhiGo/others"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Binary_Sum.go")
	fmt.Println("2. Run Climb_Stairs.go")
	fmt.Println("3. Run Square_Root.go")
	fmt.Println("4. Run Valid_Parenthesis.go")
	fmt.Println("5. Run Pascals_Triangle.go")
	fmt.Println("6. Run Buy_and_Sell_Stock.go")
	fmt.Println("7. Run Is_Palindrome.go")
	fmt.Println("8. Run Covert_excel_column_to_Alphabet_Name.go")
	fmt.Println("9. Run Excel_Title_to_Number.go")
	fmt.Println("10. Run Reverse_Bits.go")
	fmt.Println("11. Run Hamming_Weight.go")
	fmt.Println("12. Run Gas_Station.go")
	fmt.Println("13. Run Zigzag_Conversion.go")

	fmt.Print("Option: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Binary_Sum.go")
		fmt.Println("=====================")
		if number == 1 {
			others.Run_Binary_Sum()
		}
		fmt.Println("=====================")
	case 2:
		fmt.Println("Hence running Climb_Stairs.go")
		fmt.Println("=====================")
		if number == 2 {
			others.Run_Climb_Stairs()
		}
		fmt.Println("=====================")
	case 3:
		fmt.Println("Hence running Square_Root.go")
		fmt.Println("=====================")
		if number == 3 {
			others.Run_Square_Root()
		}
		fmt.Println("=====================")
	case 4:
		fmt.Println("Hence running Valid_Parenthesis.go")
		fmt.Println("=====================")
		if number == 4 {
			others.Run_Valid_Parenthesis()
		}
		fmt.Println("=====================")
	case 5:
		fmt.Println("Hence running Pascals_Triangle.go")
		fmt.Println("=====================")
		if number == 5 {
			others.Run_Pascals_Triangle()
		}
		fmt.Println("=====================")
	case 6:
		fmt.Println("Hence running Buy_and_Sell_Stock.go")
		fmt.Println("=====================")
		if number == 6 {
			others.Run_Buy_and_Sell_Stock()
		}
		fmt.Println("=====================")
	case 7:
		fmt.Println("Hence running Is_Palindrome.go")
		fmt.Println("=====================")
		if number == 7 {
			others.Run_Is_Palindrome()
		}
		fmt.Println("=====================")
	case 8:
		fmt.Println("Hence running Covert_excel_column_to_Alphabet_Name.go")
		fmt.Println("=====================")
		if number == 8 {
			others.Run_Covert_excel_column_to_Alphabet_Name()
		}
		fmt.Println("=====================")
	case 9:
		fmt.Println("Hence running Excel_Title_to_Number.go")
		fmt.Println("=====================")
		if number == 9 {
			others.Run_Excel_Title_to_Number()
		}
		fmt.Println("=====================")
	case 10:
		fmt.Println("Hence running Reverse_Bits.go")
		fmt.Println("=====================")
		if number == 10 {
			others.Run_Reverse_Bits()
		}
		fmt.Println("=====================")
	case 11:
		fmt.Println("Hence running Hamming_Weight.go")
		fmt.Println("=====================")
		if number == 11 {
			others.Run_Hamming_Weight()
		}
		fmt.Println("=====================")
	case 12:
		fmt.Println("Hence running Gas_Station.go")
		fmt.Println("=====================")
		if number == 12 {
			others.Run_Gas_Station()
		}
		fmt.Println("=====================")
	case 13:
		fmt.Println("Hence running Zigzag_Conversion.go")
		fmt.Println("=====================")
		if number == 13 {
			others.Run_Zigzag_Conversion()
		}
		fmt.Println("=====================")
	default:
		fmt.Println("Please enter valid option.")
	}
}
