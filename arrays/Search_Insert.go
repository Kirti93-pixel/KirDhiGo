package arrays

import "fmt"

//Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//You must write an algorithm with O(log n) runtime complexity.

func SearchInsertOpt(nums []int, target int) int {
	j := 0
	for _, v := range nums {
		if v < target {
			j++
		}
	}
	return j
}

func SearchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	mid := -1
	for min <= max {
		mid = (min + max) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			max = mid - 1
		} else if target > nums[mid] {
			min = mid + 1
		}
	}
	return min
}

func Run_Search_Insert() {
	fmt.Println("Index is:::", SearchInsert([]int{1, 2, 4, 5, 6, 7}, 4))
	fmt.Println("Index with Optimized Logic is:::", SearchInsertOpt([]int{1, 2, 4, 5, 6, 7}, 4))
}
