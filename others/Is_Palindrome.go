package others

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphaNumRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func isPalindrome(s string) bool {
   if len(s) == 0 {
       return true
   }
   s = strings.ToLower(s)
   s = nonAlphaNumRegex.ReplaceAllString(s, "")
   if len(s) == 1 {
       return true
   }
   lenS := len(s)
   first, last := 0, lenS-1
   for first < last {
       if s[first] != s[last] {
           return false
       }
       first++
       last--
   }
    return true
}

func Run_Is_Palindrome() {
	fmt.Println("Is string plaindrome \"A man, a plan, a canal: Panama\" ::", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("Is string plaindrome \"racing car 09\" ::", isPalindrome("racing car 09"))
}