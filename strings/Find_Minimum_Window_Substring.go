package strings

import "fmt"

func minWindow(s string, t string) string {
    answer, n, m, targetMap, tRune, sRune, matchCount, start, minLength, minStart := "", len(t), len(s), make(map[rune]int, len(t)),[]rune(t),[]rune(s), 0, 0, len(s)+2, 0

    if n > m {
        return answer
    }

    for _, r := range tRune {
        count, _ := targetMap[r]
        targetMap[r] = count+1
    }

    for i, r := range sRune {
        if count, okay := targetMap[r]; okay {
            if count>0 {
                matchCount++
            }
            targetMap[r] = count-1
        }
        if matchCount == n {
            tempCount, okay := targetMap[sRune[start]]
            for (!okay && start < i) || (okay && tempCount < 0) {
                if okay {
                    targetMap[sRune[start]] = tempCount + 1
                }
                start++
                tempCount, okay = targetMap[sRune[start]]
            }
            if minLength > i-start+1 {
                answer = string(sRune[start:i+1])
                minStart = start
                minLength = i-start+1
            }
        }
    }
    if minLength == m+2 {
        return ""
    }
    answer = string(sRune[minStart:minStart+minLength])
    return answer
}

func Run_Find_Minimum_Window_Substring() {
	s, t := "ADOBECODEBANC", "ABC"
	res := minWindow(s, t)
	fmt.Println("res::", res)
}