package ll

import "fmt"

type Node struct {
	data string
	pre  *Node
	nxt  *Node
}

type LL struct {
	head  *Node
	tail  *Node
	lenLL int
}

func (l *LL) IsEmpty() bool {
	if l.head == nil && l.tail == nil && l.lenLL == 0 {
		return true
	}
	return false
}

func (l *LL) InsertAtFirst(s string) {
	if l.IsEmpty() {
		lNode := Node{
			data: s,
			pre:  nil,
			nxt:  nil,
		}
		l.head = &lNode
		l.tail = &lNode
		l.lenLL = 1
		return
	}
	lNode := Node{
		data: s,
		pre:  nil,
		nxt:  l.head,
	}
	l.head.pre = &lNode
	l.head = l.head.pre
	l.lenLL += 1
}

func (l *LL) InsertAtLast(s string) {
	if l.IsEmpty() {
		lNode := Node{
			data: s,
			pre:  nil,
			nxt:  nil,
		}
		l.head = &lNode
		l.tail = &lNode
		l.lenLL = 1
		return
	}
	lNode := Node{
		data: s,
		pre:  l.tail,
		nxt:  nil,
	}
	l.tail.nxt = &lNode
	l.tail = l.tail.nxt
	l.lenLL += 1
}

func (l *LL) InsertAtkthPosition(s string, k int) {
	if l.IsEmpty() {
		lNode := Node{
			data: s,
			pre:  nil,
			nxt:  nil,
		}
		l.head = &lNode
		l.tail = &lNode
		l.lenLL = 1
		return
	}
	if k == 1 {
		l.InsertAtFirst(s)
		return
	}
	if k == l.lenLL+1 {
		l.InsertAtLast(s)
		return
	}
	if k > l.lenLL {
		fmt.Println("Invalid position to insert")
		return
	}

	curNode := l.head
	for i := 0; i < k-1; i++ {
		curNode = curNode.nxt
	}
	lNode := Node{
		data: s,
		nxt:  curNode,
		pre:  curNode.pre,
	}
	curNode.pre.nxt = &lNode
	curNode.pre = &lNode
	l.lenLL += 1
}

func (l *LL) DeleteFromFirst() bool {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty, no elements to delete... :)")
		return false
	}
	l.head = l.head.nxt
	l.head.pre = nil
	l.lenLL -= 1
	return true
}

func (l *LL) DeleteFromLast() bool {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty, no elements to delete... :)")
		return false
	}
	l.tail = l.tail.pre
	l.tail.nxt = nil
	l.lenLL -= 1
	return true
}

func (l *LL) DeleteFromkthPosition(k int) bool {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty, no elements to delete... :)")
		return false
	}
	if l.lenLL == 1 && k == 1 {
		l.head = nil
		l.tail = nil
		l.lenLL = 0
		return true
	}
	if k == 1 {
		l.DeleteFromFirst()
		return true
	}
	if k == l.lenLL {
		l.DeleteFromLast()
		return true
	}
	if k > l.lenLL {
		fmt.Println("Invalid position to delete... :)")
		return false
	}

	curNode := l.head
	for i := 0; i < k-1; i++ {
		curNode = curNode.nxt
	}
	curNode.pre.nxt = curNode.nxt
	curNode.nxt.pre = curNode.pre
	l.lenLL -= 1
	return true
}

func (l *LL) GetAllElements() {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty, no elements present... :)")
		return
	}
	fmt.Print("GetAllElements::")
	curNode := l.head
	for curNode != nil {
		fmt.Print(curNode.data, " ")
		curNode = curNode.nxt
	}
	fmt.Println()
}

func (l *LL) GetElementAtkthPosition(k int) (string, bool) {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty, no elements present... :)")
		return "", false
	}
	if k == 1 {
		return l.head.data, true
	}
	if k == l.lenLL {
		return l.tail.data, true
	}
	if k > l.lenLL {
		return "Invalid position to retrieve... :)", false
	}
	curNode := l.head
	for i := 0; i < k-1; i++ {
		curNode = curNode.nxt
	}
	return curNode.data, true
}

func (l *LL) GetLength() int {
	return l.lenLL
}

func Run_Linked_List() {
	var myl LL
	myl.InsertAtFirst("1")
	myl.GetAllElements()
	myl.InsertAtkthPosition("67", 1)
	myl.GetAllElements()
	elm, ok := myl.GetElementAtkthPosition(2)
	fmt.Println("elm:", elm, "ok:", ok)
	myl.InsertAtkthPosition("68", 2)
	myl.GetAllElements()
	myl.InsertAtkthPosition("69", 2)
	myl.GetAllElements()
	myl.DeleteFromkthPosition(4)
	myl.GetAllElements()
	fmt.Println("Length::", myl.GetLength())
	myl.InsertAtFirst("2")
	myl.InsertAtLast("3")
	myl.GetAllElements()
	myl.InsertAtkthPosition("67", 2)
	myl.GetAllElements()
	myl.DeleteFromFirst()
	myl.GetAllElements()
	myl.DeleteFromLast()
	myl.GetAllElements()
	myl.InsertAtFirst("56")
	myl.InsertAtFirst("29")
	myl.InsertAtFirst("50")
	myl.InsertAtFirst("20")
	myl.InsertAtFirst("11")
	myl.InsertAtLast("87")
	myl.GetAllElements()
	myl.DeleteFromkthPosition(5)
	myl.GetAllElements()
	elm, ok = myl.GetElementAtkthPosition(4)
	fmt.Println("elm:", elm, "ok:", ok)
	elm, ok = myl.GetElementAtkthPosition(5)
	fmt.Println("elm:", elm, "ok:", ok)
	elm, ok = myl.GetElementAtkthPosition(9)
	fmt.Println("elm:", elm, "ok:", ok)
	fmt.Println("Length::", myl.GetLength())

}
