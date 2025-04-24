package arrays

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	// Transform the original string to handle even-length palindromes
	// Example: "abba" => "^#a#b#b#a#$"
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)

	P := make([]int, n) // P[i] stores palindrome radius around i
	C, R := 0, 0        // C = center of current palindrome, R = right edge

	for i := 1; i < n-1; i++ {
		mirror := 2*C - i // Mirror index of i w.r.t. center C

		// If i is within the known palindrome boundary R
		if R > i {
			P[i] = min(R-i, P[mirror])
		}

		// Try to expand palindrome centered at i
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		// If the new palindrome goes beyond R, update center and right edge
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	// Find the longest palindrome in P[]
	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}

	// Extract the substring from original string s
	start := (centerIndex - maxLen) / 2 // Adjust because T has extra characters
	return s[start : start+maxLen]
}

// Utility function to get min of two ints
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Run_Longest_Palindrome() {
	fmt.Println("Longest Palindrome in string \"babad\" is", longestPalindrome("babad"))
	fmt.Println("Longest Palindrome in string \"cbbd\" is", longestPalindrome("cbbd"))
}
