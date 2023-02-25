package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/queue"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Queue.go")
	fmt.Print("Option: ")
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