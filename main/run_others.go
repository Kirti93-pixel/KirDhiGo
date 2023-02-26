package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/others"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Binary_Sum.go")
	fmt.Println("2. Run Climb_Stairs.go")
	fmt.Println("3. Run Square_Root.go")
	fmt.Println("4. Run Valid_Parenthesis.go")
	fmt.Println("5. Run Pascals_Triangle.go")

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
	default:
		fmt.Println("Please enter valid option.")
	}
}