package main

import "fmt"

func climbStairs(n int) int {
	//Its the simple fibonnaci series
	if n <= 2 {
		return n
	}
	ways, first, second := 0, 1, 2
	for i := 2; i < n; i++ {
		ways = first + second
		first, second = second, ways
	}
	return ways
}

func main() {
	fmt.Println("Number of ways to climb 5th stair::", climbStairs(5))
	fmt.Println("Number of ways to climb 3rd stair::", climbStairs(3))
	fmt.Println("Number of ways to climb 18th stair::", climbStairs(18))
}
