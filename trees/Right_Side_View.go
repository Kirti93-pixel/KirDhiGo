package trees

import "fmt"

func rightSideView(root *NodeValueString) []string {
    if root == nil {return nil}
    levels := make([]string, 0)
    calRes(root, &levels, 0)
    return levels
}

func calRes(root *NodeValueString, levels *[]string, depth int) {
    if root == nil {
        return
    }

    if len(*levels) < depth+1 {
        *levels = append(*levels, root.value)
    }
    calRes(root.right, levels, depth+1)
    calRes(root.left, levels, depth+1)
}

func Run_Right_Side_View() {
	var root *NodeValueString
	root = InsertBST2(root, "7")
	root = InsertBST2(root, "4")
	root = InsertBST2(root, "3")
	root = InsertBST2(root, "1")
	root = InsertBST2(root, "8")
	root = InsertBST2(root, "9")
	root = InsertBST2(root, "5")
	rightSideViewRes := rightSideView(root)
	fmt.Println("Right Side View:::", rightSideViewRes)
}