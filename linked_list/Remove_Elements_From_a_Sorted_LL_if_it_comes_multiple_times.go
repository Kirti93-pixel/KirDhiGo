package ll

import "fmt"

func deleteDuplicatesIfComesMultipleTimes(head *ListNode) *ListNode {
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

func Run_Remove_Elements_From_a_Sorted_LL_if_it_comes_multiple_times() {
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
	printElements(deleteDuplicatesIfComesMultipleTimes(list))
	fmt.Println()
}
