
package trees

import "fmt"

func isSameTree(p *NodeDataByte, q *NodeDataByte) bool {
    if p == nil || q == nil {
        return p == nil && q == nil
    }
    return string(p.key) == string(q.key) && isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}

func Run_Two_BSTs_Identical_or_not() {
	t1,t2 := &TreeNodeDataByte{}, &TreeNodeDataByte{}
	t1.InsertBST('F')
	t1.InsertBST('B')
	t1.InsertBST('A')
	t1.InsertBST('D')
	t1.InsertBST('C')
	t1.InsertBST('E')
	t1.InsertBST('G')
	t1.InsertBST('I')
	t1.InsertBST('H')

	t2.InsertBST('F')
	t2.InsertBST('B')
	t2.InsertBST('A')
	t2.InsertBST('D')
	t2.InsertBST('C')
	t2.InsertBST('E')
	t2.InsertBST('G')
	t2.InsertBST('I')
	t2.InsertBST('H')

	fmt.Println("Are Two trees same:::", isSameTree(t1.root,t2.root))
}