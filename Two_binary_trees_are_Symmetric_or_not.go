
package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

type Queue []*Node

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node, true
}

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insertBinaryTree(data)
	}
}

func (n *Node) insertBinaryTree(data byte) { //Level order transversal insertion
	var queue Queue
	queue.Enqueue(n)

	for !queue.IsEmpty() {
		tmp, isTmp := queue.Dequeue()
		if isTmp {
			if tmp.left != nil {
				queue.Enqueue(tmp.left)
			} else {
				tmp.left = &Node{
					key:data,
				}
				return
			}

			if tmp.right != nil {
				queue.Enqueue(tmp.right)
			} else {
				tmp.right = &Node{
					key:data,
				}
				return
			}
		}
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Println(string(n.key))
		printInOrder(n.right)
	}
}

func isSymmetric(root *Node) bool {
    return mirrorCheck(root.left, root.right)
}

func mirrorCheck(p,q *Node) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil || string(p.key) != string(q.key) {
        return false
    }
    return mirrorCheck(p.left, q.right) && mirrorCheck(p.right, q.left)
}

func main() {
	t := &Tree{}
	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')

	fmt.Println("Are Two trees symmetric:::", isSymmetric(t.root))

	//Test Binary Tree insertion
	// t1 := &Tree{
	// 	root: &Node{
	// 		key:'R',
	// 	},
	// }
	// t1.root.left = &Node{
	// 	key:'A',
	// }
	// t1.root.left.left = &Node{
	// 	key: 'B',
	// }
	// t1.root.right = &Node{
	// 	key: 'C',
	// }
	// t1.root.right.left = &Node{
	// 	key: 'D',
	// }
	// t1.root.right.right = &Node{
	// 	key: 'E',
	// }
	// fmt.Println("Inorder tranversal before insertion::")
	// printInOrder(t1.root)
	// t1.insert('H')
	// fmt.Println("Inorder tranversal afterinsertion::")
	// printInOrder(t1.root)
}