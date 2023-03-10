package ll

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var finalList *ListNode
	tmp1, tmp2 := list1, list2

	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			finalList = addNode(finalList, tmp1.Val)
			tmp1 = tmp1.Next
		} else {
			finalList = addNode(finalList, tmp2.Val)
			tmp2 = tmp2.Next
		}
	}
	tmp := finalList
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	if tmp1 == nil {
		tmp.Next = tmp2
	}
	if tmp2 == nil {
		tmp.Next = tmp1
	}
	return finalList
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

func Run_Merge_2_sorted_LL() {
	var list1, list2 *ListNode
	list1 = addNode(list1, 1)
	list1 = addNode(list1, 2)
	list1 = addNode(list1, 4)
	list1 = addNode(list1, 6)
	list1 = addNode(list1, 9)

	list2 = addNode(list2, 3)
	list2 = addNode(list2, 5)
	list2 = addNode(list2, 7)
	list2 = addNode(list2, 8)
	list2 = addNode(list2, 90)
	list2 = addNode(list2, 230)
	fmt.Print("After merging 1,2,4,6,9 and 3,5,7,8,90,230 ===>>> ")
	printElements(mergeTwoLists(list1, list2))
	fmt.Println()
}