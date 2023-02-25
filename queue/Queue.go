package queue

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(s string) {
	*q = append(*q, s)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	elm := (*q)[0]
	*q = (*q)[1:]
	return elm, true
}

func Run_Queue() {
	var myq Queue
	myq.Enqueue("1")
	myq.Enqueue("2")
	myq.Enqueue("3")

	for !myq.IsEmpty() {
		val, ok := myq.Dequeue()
		fmt.Println("val:", val, "ok:", ok)
	}
}

