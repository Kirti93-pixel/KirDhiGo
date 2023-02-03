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

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Println(string(n.key))
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Println(string(n.key))
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
	fmt.Println("PreOrder:::")
	printPreOrder(t.root)
	fmt.Println("PostOrder:::")
	printPostOrder(t.root)
	fmt.Println("InOrder:::")
	printInOrder(t.root)

}