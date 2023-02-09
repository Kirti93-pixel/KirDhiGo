package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmp := head
	for tmp != nil {
		for tmp.Next != nil && tmp.Next.Val == tmp.Val {
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return head
}

func addNode(ln *ListNode, val int) *ListNode {
	if ln == nil {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}
	tmp := ln
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	return ln
}

func printElements(list *ListNode) {
	for list != nil {
		if list.Next == nil {
			fmt.Print(list.Val)
		} else {
			fmt.Print(list.Val, ",")
		}
		list = list.Next
	}
}

func main() {
	var list *ListNode
	list = addNode(list, 1)
	list = addNode(list, 1)
	list = addNode(list, 2)
	list = addNode(list, 2)
	list = addNode(list, 3)
	list = addNode(list, 3)
	list = addNode(list, 6)
	fmt.Print("After deleting duplicates from list 1,1,2,2,3,3,6, the list is ")
	printElements(deleteDuplicates(list))
}
