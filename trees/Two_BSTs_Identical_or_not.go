
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
		t.root.insertBST(data)
	}
}

func (n *Node) insertBST(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insertBST(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insertBST(data)
		}
	}
}

func isSameTree(p *Node, q *Node) bool {
    if p == nil || q == nil {
        return p == nil && q == nil
    }
    return string(p.key) == string(q.key) && isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}

func main() {
	t1,t2 := &Tree{}, &Tree{}
	t1.insert('F')
	t1.insert('B')
	t1.insert('A')
	t1.insert('D')
	t1.insert('C')
	t1.insert('E')
	t1.insert('G')
	t1.insert('I')
	t1.insert('H')

	t2.insert('F')
	t2.insert('B')
	t2.insert('A')
	t2.insert('D')
	t2.insert('C')
	t2.insert('E')
	t2.insert('G')
	t2.insert('I')
	t2.insert('H')

	fmt.Println("Are Two trees same:::", isSameTree(t1.root,t2.root))
}