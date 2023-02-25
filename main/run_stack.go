package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/stack"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Stack.go")
	fmt.Print("Option: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Stack.go")
		fmt.Println("=====================")
		stack.Run_Stack()
		fmt.Println("=====================")
	}
}