package trees

import (
	"fmt"
)

type NAryTreeNodeOpt struct {
	ownEmpId int
	mgrEmpId int
	sal      float32
	sub      []*NAryTreeNodeOpt
	empMap   map[int]*NAryTreeNodeOpt // Map to store references to nodes for quick lookup
}

func Run_NAry_Tree_Sal_Problem_Opt() {
	// Step 1: Create the root node (CEO)
	ceo := &NAryTreeNodeOpt{
		ownEmpId: 1,
		mgrEmpId: -1, // No manager for CEO
		sal:      50,
		sub:      make([]*NAryTreeNodeOpt, 0),
		empMap:   make(map[int]*NAryTreeNodeOpt),
	}

	// Insert CEO into the map
	ceo.empMap[ceo.ownEmpId] = ceo

	// Step 2: Build the employee hierarchy by inserting nodes
	ceo.insertNodeInEmpHierarchyOpt(2, 1, 42)
	ceo.insertNodeInEmpHierarchyOpt(3, 1, 43)
	ceo.insertNodeInEmpHierarchyOpt(4, 1, 44)
	ceo.insertNodeInEmpHierarchyOpt(5, 2, 45)
	ceo.insertNodeInEmpHierarchyOpt(6, 2, 46)
	ceo.insertNodeInEmpHierarchyOpt(7, 2, 47)
	ceo.insertNodeInEmpHierarchyOpt(8, 3, 48)
	ceo.insertNodeInEmpHierarchyOpt(9, 3, 49)
	ceo.insertNodeInEmpHierarchyOpt(10, 3, 50)
	ceo.insertNodeInEmpHierarchyOpt(11, 4, 51)
	ceo.insertNodeInEmpHierarchyOpt(12, 4, 52)
	ceo.insertNodeInEmpHierarchyOpt(13, 4, 53)
	ceo.insertNodeInEmpHierarchyOpt(14, 5, 54)
	ceo.insertNodeInEmpHierarchyOpt(15, 5, 55)
	ceo.insertNodeInEmpHierarchyOpt(16, 6, 56)
	ceo.insertNodeInEmpHierarchyOpt(17, 7, 57)
	ceo.insertNodeInEmpHierarchyOpt(18, 9, 58)
	ceo.insertNodeInEmpHierarchyOpt(19, 12, 59)

	// Step 3: Print the full tree in pre-order
	fmt.Print("Input Array = ")
	ceo.preOrderNAryOpt() // Output should be in pre-order form
	fmt.Println()

	// Step 4: Define the employee for whom we want to calculate avg salary of subordinates
	fmt.Println("2 is the empId for which we need to calculate avg sal of all subordinate employees which come under given empId...")
	searchEmpId := 2

	// Step 5: Search for that employee node
	searchedEmp := ceo.searchNodeInEmpHierarchyOpt(searchEmpId)

	// Step 6: Print the subtree rooted at the searched employee
	fmt.Print("Array for only 2 as root empId = ")
	searchedEmp.preOrderNAryOpt()
	fmt.Println()

	// Step 7: Calculate count and sum of salaries of all nodes under that employee (including the employee)
	fmt.Println("Calculating the count and sum of all subordinate nodes given the new root node...")
	t, s := searchedEmp.sumOfNodesOpt()
	fmt.Println("Total nodes =", t, " | Sum of salaries =", s)

	// Step 8: Calculate average salary
	fmt.Printf("Avg = %.2f\n", s/float32(t))
}

func (n *NAryTreeNodeOpt) insertNodeInEmpHierarchyOpt(ownEmpId, mgrEmpId int, sal float32) bool {
	// If the manager is found in the map, insert the employee under the manager
	if mgrNode, found := n.empMap[mgrEmpId]; found {
		// Manager found — create the new employee node
		newEmp := &NAryTreeNodeOpt{
			ownEmpId: ownEmpId,
			mgrEmpId: mgrEmpId,
			sal:      sal,
			sub:      make([]*NAryTreeNodeOpt, 0),
			empMap:   make(map[int]*NAryTreeNodeOpt), // Initialize a map for subordinates
		}
		// Add the new employee as a child (subordinate)
		mgrNode.sub = append(mgrNode.sub, newEmp)

		// Add new employee to the map
		mgrNode.empMap[ownEmpId] = newEmp
		return true // Insertion successful
	}

	// Step 2: If manager is not found at this node, recurse through all subordinates
	for _, s := range n.sub {
		if s.insertNodeInEmpHierarchyOpt(ownEmpId, mgrEmpId, sal) {
			return true
		}
	}

	// Step 3: If manager was not found anywhere in the tree
	return false
}

func (n *NAryTreeNodeOpt) searchNodeInEmpHierarchyOpt(ownEmpId int) *NAryTreeNodeOpt {
	// Base case: check if the node is in the map
	if emp, found := n.empMap[ownEmpId]; found {
		return emp
	}

	// Recursively search through subordinates
	for _, s := range n.sub {
		if found := s.searchNodeInEmpHierarchyOpt(ownEmpId); found != nil {
			return found
		}
	}

	// If not found, return nil
	return nil
}

func (n *NAryTreeNodeOpt) sumOfNodesOpt() (int, float32) {
	// Initialize the count and salary with the current node
	curEmpCount := 1
	curSalCount := n.sal

	// Recurse through subordinates and accumulate count and salary
	for _, s := range n.sub {
		t, sal := s.sumOfNodesOpt()
		curEmpCount += t
		curSalCount += sal
	}

	// Return the total count and salary
	return curEmpCount, curSalCount
}

func (n *NAryTreeNodeOpt) preOrderNAryOpt() {
	if n == nil {
		return
	}

	// Print current employee's ID
	fmt.Print(n.ownEmpId, " ")

	// Recursively visit all subordinates
	for _, s := range n.sub {
		s.preOrderNAryOpt()
	}
}
