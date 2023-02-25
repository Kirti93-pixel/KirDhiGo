package main

import (
	"KirDhiGo/stack"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Input a number to execute any of the below:::")
	fmt.Println("1. Run Stack.go")
	fmt.Print("number: ")
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