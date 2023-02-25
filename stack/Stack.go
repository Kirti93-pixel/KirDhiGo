package stack

import "fmt"

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

func Run_Stack() {
	var st Stack
	st.Push("1")
	st.Push("2")
	st.Push("3")

	for !st.IsEmpty() {
		val, ok := st.Pop()
		fmt.Println("val:", val, "ok:", ok)
	}
}
