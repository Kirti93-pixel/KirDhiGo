package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/trees"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run BST.go")
	fmt.Println("2. Run DFS_BFS.go")
	fmt.Println("3. Run Print_all_paths_of_a_tree.go")
	fmt.Println("4. Run Two_binary_trees_are_Symmetric_or_not.go")
	fmt.Println("5. Run Two_BSTs_Identical_or_not.go")

	fmt.Print("Option: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running BST.go")
		fmt.Println("=====================")
		if number == 1 {
			trees.Run_BST()
		}
		fmt.Println("=====================")
	case 2:
		fmt.Println("Hence running DFS_BFS.go")
		fmt.Println("=====================")
		if number == 2 {
			trees.Run_DFS_BFS()
		}
		fmt.Println("=====================")
	case 3:
		fmt.Println("Hence running Print_all_paths_of_a_tree.go")
		fmt.Println("=====================")
		if number == 3 {
			trees.Run_Print_all_paths_of_a_tree()
		}
		fmt.Println("=====================")
	case 4:
		fmt.Println("Hence running Two_binary_trees_are_Symmetric_or_not.go")
		fmt.Println("=====================")
		if number == 4 {
			trees.Run_Two_binary_trees_are_Symmetric_or_not()
		}
		fmt.Println("=====================")
	case 5:
		fmt.Println("Hence running Longest_Common_Prefix.go")
		fmt.Println("=====================")
		if number == 5 {
			trees.Run_Two_BSTs_Identical_or_not()
		}
		fmt.Println("=====================")
	}
}