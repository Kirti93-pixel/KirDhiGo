package arrays

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	al, bl := len(nums1), len(nums2)
	low, high := 0, al
	var maxLeftA, minRightA, maxLeftB, minRightB int = 0, 0, 0, 0
	for low <= high {
		paA := (low + high) / 2
		paB := (al+bl+1)/2 - paA
		if paA == 0 {
			maxLeftA = math.MinInt64
		} else {
			maxLeftA = nums1[paA-1]
		}

		if paA == al {
			minRightA = math.MaxInt64
		} else {
			minRightA = nums1[paA]
		}

		if paB == 0 {
			maxLeftB = math.MinInt64
		} else {
			maxLeftB = nums2[paB-1]
		}

		if paB == bl {
			minRightB = math.MaxInt64
		} else {
			minRightB = nums2[paB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (al+bl)%2 == 0 {
				var x, y int
				if maxLeftA >= maxLeftB {
					x = maxLeftA
				} else {
					x = maxLeftB
				}

				if minRightA <= minRightB {
					y = minRightA
				} else {
					y = minRightB
				}
				xFl, yFl := float64(x), float64(y)
				return (xFl + yFl) / 2
			} else {
				if maxLeftA >= maxLeftB {
					return float64(maxLeftA)
				} else {
					return float64(maxLeftB)
				}
			}
		} else if maxLeftA > minRightB {
			high = paA - 1
		} else {
			low = paA + 1
		}
	}
	return -1
}

func Run_Median_for_array_merged_using_different_2_sized_arrays() {
	//a := []int32{1, 3, 8, 9, 15}
	//b := []int32{7, 11, 19, 21, 22, 25}  both testcases ans 11
	//a := []int32{2, 3, 5, 8}
	//b := []int32{10, 12, 14, 16, 18, 20}
	a := []int{1, 2}
	b := []int{3, 4}
	fmt.Println("Median for 1,2 and 3,4 is",findMedianSortedArrays(a, b))
}