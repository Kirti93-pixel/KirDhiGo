package arrays

import (
	"fmt"
)

//Provided a non-empty array of integers nums, every element appears twice except for one. Find that single one. Hint: Use XOR operation.

func singleNumber(nums []int) int {
    xor := 0
    for _, num := range nums {
        xor ^= num
    }
    return xor
}

func Run_Single_Number() {
	fmt.Println("Single number in the array ",[]int{1,2,2,3,3,5,5,9,1,9,8,7,7}, "is", singleNumber([]int{1,2,2,3,3,5,5,9,1,9,8,7,7}))
	fmt.Println("Single number in the array ",[]int{8,8,2,2,6,6,9,9}, "is", singleNumber([]int{8,8,2,2,6,6,9,9}))
}