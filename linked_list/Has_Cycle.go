package ll

import "fmt"

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}

func addNodeToParticularNode(ln, lnNext *ListNode, val int) *ListNode {
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
		Next: lnNext,
	}
	return ln
}

func Run_Has_Cycle() {
	var list *ListNode
	list = addNode(list, 1)
	list = addNode(list, 2)
	list = addNode(list, 4)
	list = addNode(list, 6)
	list = addNode(list, 9)
	list = addNodeToParticularNode(list, list, 10)
	fmt.Print("Does Linked-list has cycle [1,2,4,6,9,10,1] ===>>> ")
	fmt.Println(hasCycle(list))
}