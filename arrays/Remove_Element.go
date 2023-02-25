package arrays

import "fmt"

func removeElement(nums []int, val int) int {
	lenNums := len(nums)
	i := 0
	for j := 0; j < lenNums; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func removeElementOpt(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:len(nums)]...)
			n--
		} else {
			i++
		}
	}
	return n
}

func Run_Remove_Element() {
	fmt.Println("Len after removal is:::", removeElement([]int{1, 2, 4, 5, 6, 7}, 4))
	fmt.Println("Len after removal with Optimized logic is:::", removeElementOpt([]int{1, 2, 4, 5, 6, 7}, 4))
}
