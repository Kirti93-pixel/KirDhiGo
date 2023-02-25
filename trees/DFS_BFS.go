package trees

import (
	"fmt"
	"math"
)

type NodeValueString struct {
	value string
	left  *NodeValueString
	right *NodeValueString
}

func InsertBST2(n *NodeValueString, v string) *NodeValueString {
	if n == nil {
		return &NodeValueString{v, nil, nil}
	} else if v <= n.value {
		n.left = InsertBST2(n.left, v)
	} else {
		n.right = InsertBST2(n.right, v)
	}
	return n
}

/* pre-oder DFS traversal */
func preorder(n *NodeValueString) {
	if n != nil {
		fmt.Printf(n.value + " ")
		preorder(n.left)
		preorder(n.right)
	}
}

/* in-oder DFS traversal */
func inorder(n *NodeValueString) {
	if n != nil {
		inorder(n.left)
		fmt.Printf(n.value + " ")
		inorder(n.right)
	}
}

/* post-oder DFS traversal */
func postorder(n *NodeValueString) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Printf(n.value + " ")
	}
}

func maxDepth(n *NodeValueString) int {
	if n == nil {return 0}
    leftLen := maxDepth(n.left)
    rightLen := maxDepth(n.right)
    return int(math.Max(float64(leftLen),float64(rightLen))) + 1
}

func minDepthDFS(n *NodeValueString) int {
    if n == nil {return 0}
    mini := math.MaxInt32
    dfsSearch(n, &mini, 0)
    return mini+1
}

func dfsSearch(n *NodeValueString, miniAdd *int, lvl int) {
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
func breadth(n *NodeValueString) {
	if n != nil {
		s := []*NodeValueString{n}	//its a queue
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

func Run_DFS_BFS() {
	var root *NodeValueString
	root = InsertBST2(root, "7")
	root = InsertBST2(root, "4")
	root = InsertBST2(root, "3")
	root = InsertBST2(root, "1")
	root = InsertBST2(root, "8")
	root = InsertBST2(root, "9")
	root = InsertBST2(root, "5")
	fmt.Println("Pre-order DFS traversal")
	preorder(root)
	fmt.Println("\nIn-order DFS traversal")
	inorder(root)
	fmt.Println("\nPost-order DFS traversal")
	postorder(root)
	fmt.Println("\nMaxDepth(using DFS) of a tree is", maxDepth(root))
	fmt.Println("MinDepth(using DFS) of a tree is", minDepthDFS(root))
	fmt.Println("BFS traversal")
	breadth(root)
}
