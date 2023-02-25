
package trees

import "fmt"

type Queue []*NodeDataByte

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(n *NodeDataByte) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() (*NodeDataByte, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	node := (*q)[0]
	*q = (*q)[1:]
	return node, true
}

func (t *TreeNodeDataByte) InsertBinaryTree(data byte) {
	if t.root == nil {
		t.root = &NodeDataByte{key: data}
	} else {
		t.root.insertBTree(data)
	}
}

func (n *NodeDataByte) insertBTree(data byte) { //Level order transversal insertion
	var queue Queue
	queue.Enqueue(n)

	for !queue.IsEmpty() {
		tmp, ok := queue.Dequeue()
		if ok {
			if tmp.left != nil {
				queue.Enqueue(tmp.left)
			} else {
				tmp.left = &NodeDataByte{
					key:data,
				}
				return
			}

			if tmp.right != nil {
				queue.Enqueue(tmp.right)
			} else {
				tmp.right = &NodeDataByte{
					key:data,
				}
				return
			}
		}
	}
}

func isSymmetric(root *NodeDataByte) bool {
    return mirrorCheck(root.left, root.right)
}

func mirrorCheck(p,q *NodeDataByte) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil || string(p.key) != string(q.key) {
        return false
    }
    return mirrorCheck(p.left, q.right) && mirrorCheck(p.right, q.left)
}

//Find minDepth of a tree using BFS
func minDepthBFS(root *NodeDataByte) int {
    //Using BFS approach
    if root == nil {
        return 0
    }
    var queue Queue
	queue.Enqueue(root)
    lvl := 1
    for !queue.IsEmpty() {
        size := len(queue)
        for i:=0; i<size; i++ {
            tmp, ok := queue.Dequeue()
            if ok {
                if tmp.left == nil && tmp.right == nil {
                    return lvl
                }
                if tmp.left != nil {
                    queue.Enqueue(tmp.left)
                }
                if tmp.right != nil {
                    queue.Enqueue(tmp.right)
                }
            }
        }
        lvl++
    }
    return lvl
}

func Run_Two_binary_trees_are_Symmetric_or_not() {
	t := &TreeNodeDataByte{}
	t.InsertBinaryTree('F')
	t.InsertBinaryTree('B')
	t.InsertBinaryTree('A')
	t.InsertBinaryTree('D')
	t.InsertBinaryTree('C')
	t.InsertBinaryTree('E')
	t.InsertBinaryTree('G')
	t.InsertBinaryTree('I')
	t.InsertBinaryTree('H')

	fmt.Println("Are Two trees symmetric:::", isSymmetric(t.root))
	fmt.Println("MinDepth(using BFS) of a tree is", minDepthBFS(t.root))

	//Test Binary Tree insertion
	// t1 := &Tree{
	// 	root: &Node{
	// 		key:'R',
	// 	},
	// }
	// t1.root.left = &Node{
	// 	key:'A',
	// }
	// t1.root.left.left = &Node{
	// 	key: 'B',
	// }
	// t1.root.right = &Node{
	// 	key: 'C',
	// }
	// t1.root.right.left = &Node{
	// 	key: 'D',
	// }
	// t1.root.right.right = &Node{
	// 	key: 'E',
	// }
	// fmt.Println("Inorder tranversal before insertion::")
	// printInOrder(t1.root)
	// t1.insert('H')
	// fmt.Println("Inorder tranversal afterinsertion::")
	// printInOrder(t1.root)
}