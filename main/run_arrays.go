package main

import (
	"fmt"

	"github.com/Kirti93-pixel/KirDhiGo/arrays"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Binary_Search.go")
	fmt.Println("2. Kadane_Algorithm_Maximum_sum_subarray.go")
	fmt.Println("3. Median_for_array_merged_using_different_2_sized_arrays.go")
	fmt.Println("4. Merge_2_sorted_array_in_1st_array.go")
	fmt.Println("5. Merge_Sort.go")
	fmt.Println("6. Plus_One_at_the_end_of_the_digit_in_Int_array.go")
	fmt.Println("7. Remove_Duplicates_From_Array.go")
	fmt.Println("8. Remove_Element.go")
	fmt.Println("9. Search_Insert.go")
	fmt.Println("10. Single_Number.go")
	fmt.Println("11. Majority_Element.go")
	fmt.Println("12. Set_Matrix_Zeroes.go")
	fmt.Println("13. Summary_Ranges.go")

	fmt.Print("Option: ")
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
	case 10:
		fmt.Println("Hence running Single_Number.go")
		fmt.Println("=====================")
		if number == 10 {
			arrays.Run_Single_Number()
		}
		fmt.Println("=====================")
	case 11:
		fmt.Println("Hence running Majority_Element.go")
		fmt.Println("=====================")
		if number == 11 {
			arrays.Run_Majority_Element()
		}
		fmt.Println("=====================")
	case 12:
		fmt.Println("Hence running Set_Matrix_Zeroes.go")
		fmt.Println("=====================")
		if number == 12 {
			arrays.Run_Set_Matrix_Zeroes()
		}
		fmt.Println("=====================")
	case 13:
		fmt.Println("Hence running Summary_Ranges.go")
		fmt.Println("=====================")
		if number == 13 {
			arrays.Run_Summary_Ranges()
		}
		fmt.Println("=====================")
	default:
		fmt.Println("Please enter valid option.")
	}
}
