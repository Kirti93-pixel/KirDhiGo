package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(v string) {
	if t.root == nil {
		t.root = &Node{data: v}
	} else {
		t.root.Insert(v)
	}
}

func (n *Node) Insert(v string) {
	if v <= n.data {
		if n.left == nil {
			n.left = &Node{data: v}
		} else {
			n.left.Insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: v}
		} else {
			n.right.Insert(v)
		}
	}
}

func (n *Node) GetAllPaths(path []string, res *[][]string) {
	if n == nil {
		return
	}
	path = append(path, n.data)
	if n.left == nil && n.right == nil {
		*res = append(*res, path)
	} else {
		n.left.GetAllPaths(path, res)
		n.right.GetAllPaths(path, res)
	}
}

func main() {
	t := &Tree{}
	t.Insert("5")
	t.Insert("3")
	t.Insert("2")
	t.Insert("4")
	t.Insert("7")
	t.Insert("6")
	t.Insert("9")
	path := make([]string, 0)
	res := make([][]string, 0)
	t.root.GetAllPaths(path, &res)
	fmt.Println("All paths::", res)

}