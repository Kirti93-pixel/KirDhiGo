package ll

import "fmt"

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

func Run_Remove_Duplicates_From_Sorted_LL() {
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
	fmt.Println()
}
