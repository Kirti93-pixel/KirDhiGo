package arrays

import "fmt"

func mergeInplace(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
	return nums1
}

func Run_Merge_2_sorted_array_in_1st_array() {
	fmt.Println("After merging 2 sorted arrays(1st array==> [1,3,5,0,0,0] and 2nd array==> [2,4,8]) in first array::::", mergeInplace([]int{1, 3, 5, 0, 0, 0}, 3, []int{2, 4, 8}, 3))
}
