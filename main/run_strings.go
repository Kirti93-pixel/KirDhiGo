package main

import (
	"github.com/Kirti93-pixel/KirDhiGo/strings"
	"fmt"
)

func main() {
	number := 0
	fmt.Println("Select any one option to execute any of the below:::")
	fmt.Println("1. Run Find_all_palindrome_partitions_in_a_string.go")
	fmt.Println("2. Run KMP_Algorithm_String_Matching_algo.go")
	fmt.Println("3. Run Length_of_the_last_word.go")
	fmt.Println("4. Run Longest_Common_Prefix.go")
	fmt.Println("5. Run Find_all_substrings_of_a_string.go")
	fmt.Println("6. Run Find_all_contiguous_substrings_of_a_string.go")
	fmt.Println("7. Run Find_Substring_with_Concatenation_of_All_Words.go")
	fmt.Println("8. Run Find_Minimum_Window_Substring.go")

	fmt.Print("Option: ")
	fmt.Scanln(&number)
	fmt.Printf("Input number is %d \n", number)
	switch number {
	case 1:
		fmt.Println("Hence running Find_all_palindrome_partitions_in_a_string.go")
		fmt.Println("=====================")
		if number == 1 {
			strings.Run_Find_all_palindrome_partitions_in_a_string()
		}
		fmt.Println("=====================")
	case 2:
		fmt.Println("Hence running KMP_Algorithm_String_Matching_algo.go")
		fmt.Println("=====================")
		if number == 2 {
			strings.Run_KMP_Algorithm_String_Matching_algo()
		}
		fmt.Println("=====================")
	case 3:
		fmt.Println("Hence running Length_of_the_last_word.go")
		fmt.Println("=====================")
		if number == 3 {
			strings.Run_Length_of_the_last_word()
		}
		fmt.Println("=====================")
	case 4:
		fmt.Println("Hence running Longest_Common_Prefix.go")
		fmt.Println("=====================")
		if number == 4 {
			strings.Run_Longest_Common_Prefix()
		}
		fmt.Println("=====================")
	case 5:
		fmt.Println("Hence running Find_all_substrings_of_a_string.go")
		fmt.Println("=====================")
		if number == 5 {
			strings.Run_Find_all_substrings_of_a_string()
		}
		fmt.Println("=====================")	
	case 6:
		fmt.Println("Hence running Run_Find_all_contiguous_substrings_of_a_string.go")
		fmt.Println("=====================")
		if number == 6 {
			strings.Run_Find_all_contiguous_substrings_of_a_string()
		}
		fmt.Println("=====================")	
	case 7:
		fmt.Println("Hence running Run_Find_Substring_with_Concatenation_of_All_Words.go")
		fmt.Println("=====================")
		if number == 7 {
			strings.Run_Find_Substring_with_Concatenation_of_All_Words()
		}
		fmt.Println("=====================")
	case 8:
		fmt.Println("Hence running Run_Find_Minimum_Window_Substring.go")
		fmt.Println("=====================")
		if number == 8 {
			strings.Run_Find_Minimum_Window_Substring()
		}
		fmt.Println("=====================")	
	default:
		fmt.Println("Please enter valid option.")
	}
}