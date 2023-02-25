package arrays

import (
	"fmt"
	"sort"
)

func Run_Binary_Search() {
	var arr []int = []int{23, 56, 11, 85, 2, 21, 12, 90, 45, 43, 87, 62, 10}
	num := 45
	sort.Ints(arr)
	fmt.Println("Arr::", arr)
	min := 0
	max := len(arr) - 1
	mid := -1
	for min <= max {
		mid = (min + max) / 2
		if num < arr[mid] {
			max = mid - 1
		} else if num > arr[mid] {
			min = mid + 1
		} else {
			break
		}
	}
	fmt.Println("found element 45 at position::", mid)
}
