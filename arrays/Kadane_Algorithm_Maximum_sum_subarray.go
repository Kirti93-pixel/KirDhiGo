package arrays

import "fmt"

func kadane(arr []int) int {
	n, sum, max, maxNum := len(arr), 0, 0, arr[0]
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		if sum > 0 {
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
		}
		
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}
	if maxNum < 0 { //if all are non-positive numbers
		max = maxNum
	}

	return max
}

func Run_Kadane_Algorithm_Maximum_sum_subarray() {
	fmt.Println("Kadane's Algorithm::: Maximum sum subarray problem")
	var arr []int = []int{-3, 4, 5, -2, -1, 8, -4}
	fmt.Println("Maximum sum:::", kadane(arr))
}