package trees

import (
	"fmt"
)

type NodeValueInt struct {
	value int
	left  *NodeValueInt
	right *NodeValueInt
}

type TreeNodeValueInt struct {
	root *NodeValueInt
}

type QueueInt []*NodeValueInt

func (q *QueueInt) IsEmpty() bool {
	return len(*q) == 0
}

func (q *QueueInt) Enqueue(n *NodeValueInt) {
	*q = append(*q, n)
}

func (q *QueueInt) Dequeue() (*NodeValueInt, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node, true
}

func (t *TreeNodeValueInt) InsertBinaryTree1(data int) {
	if t.root == nil {
		t.root = &NodeValueInt{value: data}
	} else {
		t.root.insertBTree1(data)
	}
}

func (n *NodeValueInt) insertBTree1(data int) { //Level order transversal insertion
	var queue QueueInt
	queue.Enqueue(n)

	for !queue.IsEmpty() {
		tmp, ok := queue.Dequeue()
		if ok {
			if tmp.left != nil {
				queue.Enqueue(tmp.left)
			} else {
				tmp.left = &NodeValueInt{
					value:data,
				}
				return
			}

			if tmp.right != nil {
				queue.Enqueue(tmp.right)
			} else {
				tmp.right = &NodeValueInt{
					value:data,
				}
				return
			}
		}
	}
}

func hasPathSum(root *NodeValueInt, targetSum int) bool {
    sum := 0
    return transverseTreeForSum(root, targetSum, sum)
}

func transverseTreeForSum(root *NodeValueInt, targetSum, sum int) bool {
    if root == nil {return false}

    if root.left == nil && root.right == nil {
        sum = sum + root.value
        if sum == targetSum {
            return true
        }
    }
    return transverseTreeForSum(root.left, targetSum, sum+root.value) || transverseTreeForSum(root.right, targetSum, sum+root.value)
}

func Run_Has_path_sum() {
	t := &TreeNodeValueInt{}
	t.InsertBinaryTree1(7)
	t.InsertBinaryTree1(5)
	t.InsertBinaryTree1(2)
	t.InsertBinaryTree1(9)
	t.InsertBinaryTree1(10)
	t.InsertBinaryTree1(11)
	t.InsertBinaryTree1(17)
	t.InsertBinaryTree1(3)
	t.InsertBinaryTree1(1)
	fmt.Println("Does sum 26 exist in any of the path of the tree::", hasPathSum(t.root, 26))
	fmt.Println("Does sum 27 exist in any of the path of the tree::", hasPathSum(t.root, 27))
}