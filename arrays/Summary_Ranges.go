package arrays

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	lenNums, start, rI := len(nums), 0, 0

	res := make([]string, 0)
	for rI < lenNums {
		if (rI+1) < lenNums && nums[rI]+1 == nums[rI+1] {
			rI++
		} else {
			startStr := strconv.Itoa(nums[start])
			rangeIdentifierStr := strconv.Itoa(nums[rI])
			if start == rI {
				res = append(res, startStr)
			} else {
				res = append(res, startStr+"->"+rangeIdentifierStr)
			}
			start = rI + 1
			rI++
		}
	}
	return res
}

func Run_Summary_Ranges() {
	input := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("1. For Input: ", input, "output is::", summaryRanges(input))
	input = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println("2. For Input: ", input, "output is::", summaryRanges(input))
}
