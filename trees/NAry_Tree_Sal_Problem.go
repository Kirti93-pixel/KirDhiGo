package trees

import (
	"fmt"
)

type NAryTreeNode struct {
	ownEmpId int
	mgrEmpId int
	sal      float32
	sub      []*NAryTreeNode
}

func Run_NAry_Tree_Sal_Problem() {
	// Step 1: Create the root node (CEO)
	ceo := &NAryTreeNode{
		ownEmpId: 1,
		mgrEmpId: -1, // No manager for CEO
		sal:      50,
		sub:      make([]*NAryTreeNode, 0),
	}

	// Step 2: Build the employee hierarchy by inserting nodes
	ceo.insertNodeInEmpHierarchy(2, 1, 42)
	ceo.insertNodeInEmpHierarchy(3, 1, 43)
	ceo.insertNodeInEmpHierarchy(4, 1, 44)
	ceo.insertNodeInEmpHierarchy(5, 2, 45)
	ceo.insertNodeInEmpHierarchy(6, 2, 46)
	ceo.insertNodeInEmpHierarchy(7, 2, 47)
	ceo.insertNodeInEmpHierarchy(8, 3, 48)
	ceo.insertNodeInEmpHierarchy(9, 3, 49)
	ceo.insertNodeInEmpHierarchy(10, 3, 50)
	ceo.insertNodeInEmpHierarchy(11, 4, 51)
	ceo.insertNodeInEmpHierarchy(12, 4, 52)
	ceo.insertNodeInEmpHierarchy(13, 4, 53)
	ceo.insertNodeInEmpHierarchy(14, 5, 54)
	ceo.insertNodeInEmpHierarchy(15, 5, 55)
	ceo.insertNodeInEmpHierarchy(16, 6, 56)
	ceo.insertNodeInEmpHierarchy(17, 7, 57)
	ceo.insertNodeInEmpHierarchy(18, 9, 58)
	ceo.insertNodeInEmpHierarchy(19, 12, 59)

	// Step 3: Print the full tree in pre-order
	fmt.Print("Input Array = ")
	ceo.preOrderNAry() // Output should be in pre-order form
	fmt.Println()

	// Step 4: Define the employee for whom we want to calculate avg salary of subordinates
	fmt.Println("2 is the empId for which we need to calculate avg sal of all subordinate employees which come under given empId...")
	searchEmpId := 2

	// Step 5: Search for that employee node
	searchedEmp := ceo.searchNodeInEmpHierarchy(searchEmpId)

	// Step 6: Print the subtree rooted at the searched employee
	fmt.Print("Array for only 2 as root empId = ")
	searchedEmp.preOrderNAry()
	fmt.Println()

	// Step 7: Calculate count and sum of salaries of all nodes under that employee (including the employee)
	fmt.Println("Calculating the count and sum of all subordinate nodes given the new root node...")
	t, s := searchedEmp.sumOfNodes()
	fmt.Println("Total nodes =", t, " | Sum of salaries =", s)

	// Step 8: Calculate average salary
	fmt.Printf("Avg = %.2f\n", s/float32(t))
}

func (n *NAryTreeNode) insertNodeInEmpHierarchy(ownEmpId, mgrEmpId int, sal float32) bool {
	// Base case: If the current node is nil, we can't insert anything
	if n == nil {
		return false
	}

	// Step 1: Check if current node is the manager under whom we want to insert
	if n.ownEmpId == mgrEmpId {
		// Manager found — create the new employee node
		newEmp := &NAryTreeNode{
			ownEmpId: ownEmpId,
			mgrEmpId: mgrEmpId,
			sal:      sal,
		}
		// Add the new employee as a child (subordinate)
		n.sub = append(n.sub, newEmp)
		return true // Insertion successful
	}

	// Step 2: If not found at this node, search recursively in all subordinates
	for _, s := range n.sub {
		if s.insertNodeInEmpHierarchy(ownEmpId, mgrEmpId, sal) {
			return true // Inserted in one of the subtrees
		}
	}

	// Step 3: If the manager was not found in the entire tree
	return false
}

func (n *NAryTreeNode) searchNodeInEmpHierarchy(ownEmpId int) *NAryTreeNode {
	// Base case: if the current node is nil, return nil (nothing to search)
	if n == nil {
		return nil
	}

	// Step 1: If the current node matches the target employee ID, return it
	if n.ownEmpId == ownEmpId {
		return n
	}

	// Step 2: Otherwise, recursively search in all child nodes (subordinates)
	for _, s := range n.sub {
		found := s.searchNodeInEmpHierarchy(ownEmpId)
		if found != nil {
			return found // Found the node in one of the subtrees
		}
	}

	// Step 3: If the node is not found in the entire subtree, return nil
	return nil
}

func (n *NAryTreeNode) sumOfNodes() (int, float32) {
	// Step 1: Start with current employee
	curEmpCount := 1     // Count the current node
	curSalCount := n.sal // Add the current employee's salary

	// Step 2: Recurse through all subordinates and add their values
	for _, s := range n.sub {
		t, sal := s.sumOfNodes() // Recursive call to get count and salary from child
		curEmpCount += t         // Add number of employees in the subtree
		curSalCount += sal       // Add their total salary
	}

	// Step 3: Return the combined result
	return curEmpCount, curSalCount
}

func (n *NAryTreeNode) preOrderNAry() {
	if n == nil {
		return // Base case: if node is nil, stop
	}

	// Step 1: Visit the current node first (pre-order)
	fmt.Print(n.ownEmpId, " ")

	// Step 2: Recursively visit all children (from left to right)
	for _, s := range n.sub {
		s.preOrderNAry()
	}
}
