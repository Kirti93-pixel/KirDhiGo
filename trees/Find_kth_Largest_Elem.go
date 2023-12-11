package trees

import "fmt"

func kthElemFinding(k int, counter *int, root *NodeValueString, ans *string) {
	if root == nil || *counter >= k {
		return
	}

	kthElemFinding(k, counter, root.right, ans)
	*counter++
	if *counter == k {
		*ans = root.value
		return
	}
	kthElemFinding(k, counter, root.left, ans)
}

func Run_Find_kth_Largest_Elem() {
	var root *NodeValueString
	root = InsertBST2(root, "g")
	root = InsertBST2(root, "d")
	root = InsertBST2(root, "e")
	root = InsertBST2(root, "f")
	root = InsertBST2(root, "j")
	root = InsertBST2(root, "i")
	root = InsertBST2(root, "h")
	counter, ans := 0, ""
	kthElemFinding(3, &counter, root, &ans)
	fmt.Println("kth largest element:::", ans)
}