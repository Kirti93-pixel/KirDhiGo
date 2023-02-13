
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

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
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
}