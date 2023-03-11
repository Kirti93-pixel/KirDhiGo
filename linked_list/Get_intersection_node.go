package ll

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    tmpA := headA
    tmpB := headB

    for tmpA != tmpB {
        if tmpA == nil {
            tmpA = headB
        } else {
            tmpA = tmpA.Next
        }

        if tmpB == nil {
            tmpB = headA
        } else {
            tmpB = tmpB.Next
        }
    }

    return tmpA
}

func Run_Get_intersection_node() {
	var list1, tmp1 *ListNode
	var list2 *ListNode
	list1 = addNode(list1, 1)
	list1 = addNode(list1, 2)
	list1 = addNode(list1, 4)
	tmp1 = addNodeAndReturnNode(list1, 89)
	list1 = addNode(list1, 7)
	list1 = addNode(list1, 9)

	list2 = addNode(list2, 10)
	list2 = addNode(list2, 20)
	tmp2 := list2
	for tmp2.Next != nil {
		tmp2 = tmp2.Next
	}
	tmp2.Next = tmp1
	fmt.Print("Does Linked-lists has intersection point list1: [1,2,4,89,7,9] and list2: [10,20,89,7,9] ===>>> ")
	res := getIntersectionNode(list1, list2)
	fmt.Println(res.Val)
}

func addNodeAndReturnNode(ln *ListNode, val int) *ListNode {
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
	return tmp.Next
}