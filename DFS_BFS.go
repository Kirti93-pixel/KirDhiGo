package main

import (
	"fmt"
	"math"
)

type node struct {
	value string
	left  *node
	right *node
}

func insert(n *node, v string) *node {
	if n == nil {
		return &node{v, nil, nil}
	} else if v <= n.value {
		n.left = insert(n.left, v)
	} else {
		n.right = insert(n.right, v)
	}
	return n
}

/* pre-oder DFS traversal */
func preorder(n *node) {
	if n != nil {
		fmt.Printf(n.value + " ")
		preorder(n.left)
		preorder(n.right)
	}
}

/* in-oder DFS traversal */
func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Printf(n.value + " ")
		inorder(n.right)
	}
}

/* post-oder DFS traversal */
func postorder(n *node) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Printf(n.value + " ")
	}
}

func maxDepth(n *node) int {
	if n == nil {return 0}
    leftLen := maxDepth(n.left)
    rightLen := maxDepth(n.right)
    return int(math.Max(float64(leftLen),float64(rightLen))) + 1
}

func minDepth(root *node) int {
    if root == nil {return 0}
    mini := math.MaxInt32
    dfsSearch(root, &mini, 0)
    return mini+1
}

func dfsSearch(n *node, miniAdd *int, lvl int) {
    if n != nil {
        if n.left == nil && n.right == nil {
            *miniAdd = int(math.Min(float64(*miniAdd), float64(lvl)))
        } else {
            dfsSearch(n.left, miniAdd, lvl+1)
            dfsSearch(n.right, miniAdd, lvl+1)
        }
    }
}
/* breadth first traversal */
func breadth(n *node) {
	if n != nil {
		s := []*node{n}	//its a queue
		for len(s) > 0 {
			current_node := s[0]
			fmt.Printf(current_node.value + " ")
			s = s[1:]
			if current_node.left != nil {
				s = append(s, current_node.left)
			}
			if current_node.right != nil {
				s = append(s, current_node.right)
			}
		}
	}
}

func main() {
	var root *node
	root = insert(root, "7")
	root = insert(root, "4")
	root = insert(root, "3")
	root = insert(root, "1")
	root = insert(root, "8")
	root = insert(root, "9")
	root = insert(root, "5")
	fmt.Println("Pre-order DFS traversal")
	preorder(root)
	fmt.Println("\nIn-order DFS traversal")
	inorder(root)
	fmt.Println("\nPost-order DFS traversal")
	postorder(root)
	fmt.Println("\nMaxDepth(using DFS) of a tree is", maxDepth(root))
	fmt.Println("MinDepth(using DFS) of a tree is", minDepth(root))
	fmt.Println("BFS traversal")
	breadth(root)
}
