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
	fmt.Println("6. Run Has_path_sum.go")
	fmt.Println("7. Run Right_Side_View.go")
	fmt.Println("8. Run Average_of_Levels.go")
	fmt.Println("9. Run Find_kth_Largest_Elem.go")

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
	case 6:
		fmt.Println("Hence running Has_path_sum.go")
		fmt.Println("=====================")
		if number == 6 {
			trees.Run_Has_path_sum()
		}
		fmt.Println("=====================")
	case 7:
		fmt.Println("Hence running Right_Side_View.go")
		fmt.Println("=====================")
		if number == 7 {
			trees.Run_Right_Side_View()
		}
		fmt.Println("=====================")
	case 8:
		fmt.Println("Hence running Average_of_Levels.go")
		fmt.Println("=====================")
		if number == 8 {
			trees.Run_Average_of_Levels()
		}
		fmt.Println("=====================")
	case 9:
		fmt.Println("Hence running Find_kth_Largest_Elem.go")
		fmt.Println("=====================")
		if number == 9 {
			trees.Run_Find_kth_Largest_Elem()
		}
		fmt.Println("=====================")
	default:
		fmt.Println("Please enter valid option.")
	}
}
