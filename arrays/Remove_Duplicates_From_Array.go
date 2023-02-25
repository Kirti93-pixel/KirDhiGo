package main

import "fmt"

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 || numsLen == 1 {
		return numsLen
	}

	insertIdx := 1
	for i := 1; i < numsLen; i++ {
		if nums[i-1] == nums[i] {
			continue //do nothing
		} else {
			nums[insertIdx] = nums[i]
			insertIdx++
		}
	}
	return insertIdx
}

func main() {
	fmt.Println("After removing duplicates, the length is::::", removeDuplicates([]int{1, 1, 1, 2, 2, 3, 4, 5, 5, 5, 6, 7, 7, 8}))
}
