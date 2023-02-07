package main

import "fmt"

func climbStairs(n int) int {
	//Its the simple fibonnaci series
	if n <= 2 {
		return n
	}
	ways, twoBelowCurrentStep, oneBelowCurrentStep := 0, 1, 2
	for i := 2; i < n; i++ {
		ways = twoBelowCurrentStep + oneBelowCurrentStep
		twoBelowCurrentStep, oneBelowCurrentStep = oneBelowCurrentStep, ways
	}
	return ways
}

func main() {
	fmt.Println("Number of ways to climb 8th stair::", climbStairs(8))
	fmt.Println("Number of ways to climb 8th stair::", climbStairs(3))
	fmt.Println("Number of ways to climb 8th stair::", climbStairs(18))
}
