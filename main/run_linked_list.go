package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/linked_list"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Linked_List.go")
	fmt.Println("2. Run Merge_2_sorted_LL.go")
	fmt.Println("3. Run Remove_Duplicates_From_Sorted_LL.go")
	fmt.Println("4. Run Remove_Elements_From_a_Sorted_LL_if_it_comes_multiple_times.go")

	fmt.Print("Option: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Linked_List.go")
		fmt.Println("=====================")
		if number == 1 {
			ll.Run_Linked_List()
		}
		fmt.Println("=====================")
	case 2:
		fmt.Println("Hence running Merge_2_sorted_LL.go")
		fmt.Println("=====================")
		if number == 2 {
			ll.Run_Merge_2_sorted_LL()
		}
		fmt.Println("=====================")
	case 3:
		fmt.Println("Hence running Remove_Duplicates_From_Sorted_LL.go")
		fmt.Println("=====================")
		if number == 3 {
			ll.Run_Remove_Duplicates_From_Sorted_LL()
		}
		fmt.Println("=====================")
	case 4:
		fmt.Println("Hence running Remove_Elements_From_a_Sorted_LL_if_it_comes_multiple_times.go")
		fmt.Println("=====================")
		if number == 4 {
			ll.Run_Remove_Elements_From_a_Sorted_LL_if_it_comes_multiple_times()
		}
		fmt.Println("=====================")
	}
}