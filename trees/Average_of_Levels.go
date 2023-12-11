package trees

import "fmt"

func averageOfLevels(root *NodeValueInt) []float64 {
    if root == nil {return []float64{}}

    curLvlQ, res := []*NodeValueInt{root}, []float64{}

    for len(curLvlQ) != 0 {
        nextLvlQ, curLvlSum := []*NodeValueInt{}, 0.0
        for _, node := range curLvlQ {
            if node != nil {
                curLvlSum += float64(node.value)
            }

            if node.left!=nil{
                nextLvlQ = append(nextLvlQ, node.left)
            }

            if node.right!=nil{
                nextLvlQ = append(nextLvlQ, node.right)
            }
        }
        curAvg := curLvlSum/float64(len(curLvlQ))
        res = append(res, curAvg)
        curLvlQ = nextLvlQ
    }
    return res
}

func Run_Average_of_Levels() {
	t := &TreeNodeValueInt{}
	t.InsertBinaryTree1(3)
	t.InsertBinaryTree1(9)
	t.InsertBinaryTree1(20)
	t.InsertBinaryTree1(15)
	t.InsertBinaryTree1(7)
	fmt.Println("Average of levels::", averageOfLevels(t.root))
}

