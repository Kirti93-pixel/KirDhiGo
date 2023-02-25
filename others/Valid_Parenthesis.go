package others

import "fmt"

func isValid(s string) bool {
	sLen := len(s)
	myStack := make(Stack, 0)

	for i := 0; i < sLen; i++ {
		if string(s[i]) == "(" {
			myStack.Push(")")
		} else if string(s[i]) == "{" {
			myStack.Push("}")
		} else if string(s[i]) == "[" {
			myStack.Push("]")
		} else if myStack.IsEmpty() {
			return false
		} else if val, okay := myStack.Pop(); okay {
			if val != string(s[i]) {
				return false
			}
		}
	}

	return myStack.IsEmpty()
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	resStr := (*s)[index]
	*s = (*s)[:index]
	return resStr, true
}

func Run_Valid_Parenthesis() {
	fmt.Println("Is parenthesis valid {}[]() ::::", isValid("{}[]()"))
	fmt.Println("Is parenthesis valid {}[](  ::::", isValid("{}[]("))
	fmt.Println("Is parenthesis valid {}() ::::", isValid("{}()"))
}