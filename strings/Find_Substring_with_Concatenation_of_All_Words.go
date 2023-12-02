package strings

import "fmt"

func findSubstring(s string, words []string) []int {
    wordLen, totalWords, mem, temp, found, result := len(words[0]), len(words), make(map[string]int, len(words)),  make(map[string]int, len(words)), false, make([]int, 0)

    for _, str := range words {
        mem[str] +=1
    } 

    for i:=0; i+wordLen*totalWords<=len(s); {
        found, temp = true, make(map[string]int, totalWords)
        for j:=i;j<i+wordLen*totalWords;j+=wordLen{
            if _, okay := mem[ s[j:j+wordLen] ]; okay {
                temp[s[j:j+wordLen]]+=1
            } else {
                found=false
                break
            }
        }
        if found{
            for key, _ := range mem {
                if val, okay := temp[key]; !okay || val != mem[key] {
                    found = false
                    i++
                    break
                }
            }
            if found {
                result = append(result, i)
                i++
            }
        } else {
            i++
        }
    }

    return result
}

func Run_Find_Substring_with_Concatenation_of_All_Words() {
	s := "barfoofoobarthefoobarman"
	res := findSubstring(s, []string{"bar","foo","the"})
	fmt.Println("res::", res)
}