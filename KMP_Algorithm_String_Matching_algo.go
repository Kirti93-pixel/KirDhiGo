package main

import "fmt"

func KMP(text, pat string) bool {
	lps := computeTemporaryArray(pat)
	var i, j int = 0, 0
	textLen, patLen := len(text), len(pat)
	for i < textLen && j < patLen {
		if string(text[i]) == string(pat[j]) {
			i, j = i+1, j+1
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	if j == patLen {
		return true
	}
	return false
}

func computeTemporaryArray(pat string) []int {
	lps := make([]int, len(pat))
	ind := 0
	patLen := len(pat)
	for i := 1; i < patLen; {
		if string(pat[i]) == string(pat[ind]) {
			lps[i] = ind + 1
			ind++
			i++
		} else {
			if ind != 0 {
				ind = lps[ind-1]
			} else {
				lps[i] = 0
				i++
			}

		}
	}
	return lps
}

func main() {
	fmt.Println("KMP String matching algorithm")
	text := "abshdbfudmabcabydgfjrndkf"
	pat := "abcaby"
	fmt.Println(KMP(text, pat))
}