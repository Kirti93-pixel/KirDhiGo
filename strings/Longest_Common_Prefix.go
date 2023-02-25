package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	//Horizontal scanning
	// lenStr := len(strs)
	// if lenStr == 0 {
	//     return ""
	// }
	// prefix := strs[0]
	// for i:=1; i<lenStr;i++ {
	//     for strings.Index(strs[i], prefix) != 0 {
	//         prefix = prefix[0:len(prefix)-1]
	//         if prefix == "" {
	//             break
	//         }
	//     }
	// }
	// return prefix

	//Vertical Scanning
	lenStr := len(strs)
	if lenStr == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

func main() {
	var strs []string = []string{"flower", "flow", "flows"}
	fmt.Println("Longest Common Prefix is:::", longestCommonPrefix(strs))

}
