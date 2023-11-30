package strings

import "fmt"

func findAllContiguousSubstrings(s string) map[string]struct{} {
	lenS := len(s)
	res := make(map[string]struct{}, 0)
	str := string(s[0])
	res[str] = struct{}{}
	for i:=1; i<lenS;i++ {
		if s[i-1] == s[i] {
			str += string(s[i])
		} else {
			str = string(s[i])
		}
		res[str] = struct{}{}
	}
	return res
}

func Run_Find_all_contiguous_substrings_of_a_string() {
	s := "abbcdhjikkkkkaaskhngj"
	res := findAllContiguousSubstrings(s)
	fmt.Println("res::", res)
}