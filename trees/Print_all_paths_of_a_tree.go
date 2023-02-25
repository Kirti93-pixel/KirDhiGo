package trees

import "fmt"

type NodeDataString struct {
	data  string
	left  *NodeDataString
	right *NodeDataString
}

type TreeNodeDataString struct {
	root *NodeDataString
}

func (t *TreeNodeDataString) InsertBST3(v string) {
	if t.root == nil {
		t.root = &NodeDataString{data: v}
	} else {
		t.root.insert(v)
	}
}

func (n *NodeDataString) insert(v string) {
	if v <= n.data {
		if n.left == nil {
			n.left = &NodeDataString{data: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &NodeDataString{data: v}
		} else {
			n.right.insert(v)
		}
	}
}

func (n *NodeDataString) GetAllPaths(path []string, res *[][]string) {
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

func Run_Print_all_paths_of_a_tree() {
	t := &TreeNodeDataString{}
	t.InsertBST3("5")
	t.InsertBST3("3")
	t.InsertBST3("2")
	t.InsertBST3("4")
	t.InsertBST3("7")
	t.InsertBST3("6")
	t.InsertBST3("9")
	path := make([]string, 0)
	res := make([][]string, 0)
	t.root.GetAllPaths(path, &res)
	fmt.Println("All paths::", res)

}
