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
	ans := &ListNode{}
	ans.Next = head
	var ptr *ListNode = ans

	for ptr.Next != nil && ptr.Next.Next != nil {
		if ptr.Next.Val == ptr.Next.Next.Val {
			value := ptr.Next.Val
			for ptr.Next != nil && ptr.Next.Val == value {
				ptr.Next = ptr.Next.Next
			}
		} else {
			ptr = ptr.Next
		}
	}
	return ans.Next
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
	list = addNode(list, 2)
	list = addNode(list, 3)
	list = addNode(list, 3)
	list = addNode(list, 4)
	list = addNode(list, 5)
	list = addNode(list, 5)
	list = addNode(list, 6)
	fmt.Print("After deleting elements(which are duplicates) from list 1,2,3,3,4,5,5,6, the list is ")
	printElements(deleteDuplicates(list))
}
