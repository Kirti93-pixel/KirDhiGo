package strings

import "fmt"

func findAllSubstrings(s string) []string {
	lenS := len(s)
	res := make([]string, 0)
	for i:=1; i<lenS;i++ {
		for j:=0;j<lenS-i+1;j++ {
			res = append(res, s[j:j+i])
		}
	}
	return res
}

func Run_Find_all_substrings_of_a_string() {
	s := "abbcdhjiaaskhngj"
	res := findAllSubstrings(s)
	fmt.Println("res::", res)
}