package strings

import "fmt"

func findAllPalPartitions(index int, s string, path []string, res *[][]string) {
	if index == len(s) {
		*res = append(*res, path)
		//fmt.Println("res", *res)
		return
	}
	for i := index; i < len(s); i++ {
		if IsPalindrome(s, index, i) {
			path = append(path, s[index:i+1])
			//fmt.Println("path:", path)
			findAllPalPartitions(i+1, s, path, res)
			if len(path) > 0 {
				path = path[:len(path)-1]
			} else if len(path) > 1 {
				path = path[:len(path)-2]
			}
		}
	}
}

func IsPalindrome(str string, start, end int) bool {
	for start <= end {
		if str[start] != str[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}

func Run_Find_all_palindrome_partitions_in_a_string() {
	s := "aabba"
	path := make([]string, 0)
	result := make([][]string, 0)
	findAllPalPartitions(0, s, path, &result)
	fmt.Println("res::", result)
}
