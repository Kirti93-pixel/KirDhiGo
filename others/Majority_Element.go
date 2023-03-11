package others

import "fmt"

func majorityElement(nums []int) int {
    lenNums := len(nums)

    if lenNums == 1 {
        return nums[0]
    }

    count := 0
    element := nums[0]
    for i:=0; i<lenNums; i++ {
        if count == 0 {
            element = nums[i]
        }
        if element == nums[i] {
            count++
        } else {
            count--
        }
    }

    return element
}

func Run_Majority_Element() {
	fmt.Println("Majority element in array [1,1,1,2,2,2,3,2,5,2,1,1,1,1] is", majorityElement([]int{1,1,1,2,2,2,3,2,5,2,1,1,1,1}))
	fmt.Println("Majority element in array [1,1,1,2,3,3,3,3,3,3,5] is", majorityElement([]int{1,1,1,2,3,3,3,3,3,3,5}))
}