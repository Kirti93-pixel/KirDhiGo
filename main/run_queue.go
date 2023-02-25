package main

import (
	"KirDhiGo/queue"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Input a number to execute any of the below:::")
	fmt.Println("1. Run Queue.go")
	fmt.Print("number: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Binary_Search.go")
		fmt.Println("=====================")
		queue.Run_Queue()
		fmt.Println("=====================")
	}
}