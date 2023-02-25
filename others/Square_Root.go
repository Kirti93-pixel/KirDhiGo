package others

import (
	"fmt"
)

func squareRoot(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	i := 1
	for ; i <= x/i; i++ {
		if x/i == i {
			return i
		}
	}
	return i - 1
}

func squareRootBinarySearch(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	low, end, ans := 1, x, 0

	for low <= end {
		mid := int((low + end) / 2)
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			low = mid + 1
			ans = mid // Just to get the lower bound Value (if we need upper bound value then we can add this in else condition)
		} else {
			end = mid - 1
		}
	}
	return ans
}

func Run_Square_Root() {
	fmt.Println("squareRoot for 16 is:::", squareRoot(16))
	fmt.Println("squareRootBinarySearch for 80 is:::", squareRootBinarySearch(80))
}
