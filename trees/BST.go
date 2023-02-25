package trees

import (
	"fmt"
	"math"
)

type TreeNodeDataByte struct {
	root *NodeDataByte
}

type NodeDataByte struct {
	key   byte
	left  *NodeDataByte
	right *NodeDataByte
}

func (t *TreeNodeDataByte) InsertBST(data byte) {
	if t.root == nil {
		t.root = &NodeDataByte{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *NodeDataByte) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &NodeDataByte{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &NodeDataByte{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPreOrder(n *NodeDataByte) {
	if n == nil {
		return
	} else {
		fmt.Print(string(n.key))
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *NodeDataByte) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Print(string(n.key))
	}
}

func printInOrder(n *NodeDataByte) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Print(string(n.key))
		printInOrder(n.right)
	}
}

func inorderTraversal(root *NodeDataByte) []string {
    res := make([]string, 0)
    treeTransverse(root, &res)
    return res
}

func treeTransverse(node *NodeDataByte, res *[]string) {
    if node == nil {
        return
    }
    treeTransverse(node.left, res)
    *res = append(*res, string(node.key))
    treeTransverse(node.right, res)
}

//Convert Sorted array to BST:
func sortedArrayToBST(nums []byte) *NodeDataByte {
    return createBST(nums, 0, len(nums))
}

func createBST(nums []byte, start, end int) *NodeDataByte {
    if start >= end {
        return nil
    }
    return &NodeDataByte{
        key: nums[(start+end)/2], //will take upper bound, for 3.5 it will take 4
        left: createBST(nums, start, (start+end)/2),
        right: createBST(nums, ((start+end)/2)+1, end),
    }
}

//Check if the BST is balanced
func isBalanced(root *NodeDataByte) bool {
    return getHeight(root) != -1
}

func getHeight(root *NodeDataByte) int {
    if root == nil {
        return 0
    }
    leftH := getHeight(root.left)
    rightH := getHeight(root.right)
    if leftH == -1 || rightH == -1 || int(math.Abs(float64(leftH) - float64(rightH))) > 1 {
        return -1
    }
    return int(math.Max(float64(leftH), float64(rightH)))+1
}

func Run_BST() {
	t := &TreeNodeDataByte{}
	t.InsertBST('F')
	t.InsertBST('B')
	t.InsertBST('A')
	t.InsertBST('D')
	t.InsertBST('C')
	t.InsertBST('E')
	t.InsertBST('G')
	t.InsertBST('I')
	t.InsertBST('H')
	fmt.Print("PreOrder:::")
	printPreOrder(t.root)
	fmt.Println()
	fmt.Print("PostOrder:::")
	printPostOrder(t.root)
	fmt.Println()
	fmt.Print("InOrder:::")
	printInOrder(t.root)
	fmt.Println()
	fmt.Println("Stored inordered transversal in an array:::", inorderTraversal(t.root))
	fmt.Println("Is Tree balanced:", isBalanced(t.root))

	nums := []byte{'A','B','C','D','E','F','G','H','I'} 
	fmt.Print("InOrder after sorted array converted to BST:::")
	printInOrder(sortedArrayToBST(nums))
}