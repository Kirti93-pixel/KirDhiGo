package main

import (
	"KirDhiGo/arrays"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Input a number to execute any of the below:::")
	fmt.Println("1. Binary_Search.go")
	fmt.Println("2. Kadane_Algorithm_Maximum_sum_subarray.go")
	fmt.Println("3. Median_for_array_merged_using_different_2_sized_arrays.go")
	fmt.Println("4. Merge_2_sorted_array_in_1st_array.go")
	fmt.Println("5. Merge_Sort.go")
	fmt.Println("6. Plus_One_at_the_end_of_the_digit_in_Int_array.go")
	fmt.Println("7. Remove_Duplicates_From_Array.go")
	fmt.Println("8. Remove_Element.go")
	fmt.Println("9. Search_Insert.go")
	fmt.Print("number: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Binary_Search.go")
		fmt.Println("=====================")
		if number == 1 {
			arrays.Run_Binary_Search()
		}
		fmt.Println("=====================")
	case 2:
		fmt.Println("Hence running Kadane_Algorithm_Maximum_sum_subarray.go")
		fmt.Println("=====================")
		if number == 2 {
			arrays.Run_Kadane_Algorithm_Maximum_sum_subarray()
		}
		fmt.Println("=====================")
	case 3:
		fmt.Println("Hence running Median_for_array_merged_using_different_2_sized_arrays.go")
		fmt.Println("=====================")
		if number == 3 {
			arrays.Run_Median_for_array_merged_using_different_2_sized_arrays()
		}
		fmt.Println("=====================")
	case 4:
		fmt.Println("Hence running Merge_2_sorted_array_in_1st_array.go")
		fmt.Println("=====================")
		if number == 4 {
			arrays.Run_Merge_2_sorted_array_in_1st_array()
		}
		fmt.Println("=====================")
	case 5:
		fmt.Println("Hence running Merge_Sort.go")
		fmt.Println("=====================")
		if number == 5 {
			arrays.Run_Merge_Sort()
		}
		fmt.Println("=====================")
	case 6:
		fmt.Println("Hence running Plus_One_at_the_end_of_the_digit_in_Int_array.go")
		fmt.Println("=====================")
		if number == 6 {
			arrays.Run_Plus_One_at_the_end_of_the_digit_in_Int_array()
		}
		fmt.Println("=====================")
	case 7:
		fmt.Println("Hence running Remove_Duplicates_From_Array.go")
		fmt.Println("=====================")
		if number == 7 {
			arrays.Run_Remove_Duplicates_From_Array()
		}
		fmt.Println("=====================")
	case 8:
		fmt.Println("Hence running Remove_Element.go")
		fmt.Println("=====================")
		if number == 8 {
			arrays.Run_Remove_Element()
		}
		fmt.Println("=====================")
	case 9:
		fmt.Println("Hence running Search_Insert.go")
		fmt.Println("=====================")
		if number == 9 {
			arrays.Run_Search_Insert()
		}
		fmt.Println("=====================")
	}
}