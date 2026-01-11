package neetcode150

// Problem: Copy List with Random Pointer
// Each node in the linked list has two pointers:
//   - Next   → points to the next node
//   - Random → points to any node in the list or nil
// We must create a deep copy of the list.
//
// Pattern:
// Hash Map + Two Pass Linked List Traversal
//
// Intuition:
// The biggest challenge is that Random pointers can point
// forward, backward, or be nil. We cannot assign Random pointers
// correctly until *all nodes are created*.
//
// Solution Idea:
// 1. First pass: Create a copy of every node (only value copied).
//    Store mapping: original node → cloned node.
// 2. Second pass: Use the map to wire up Next and Random pointers
//    correctly for the cloned nodes.
//
// Why Hash Map?
// - It allows O(1) lookup to find the cloned version of any original node.
//
// Time Complexity: O(n)
// Space Complexity: O(n)

func copyRandomList(head *Node) *Node {
	// Map to store: original node → copied node
	m := make(map[*Node]*Node)

	// -------------------------
	// Pass 1: Clone all nodes (values only)
	// -------------------------
	current := head
	for current != nil {
		// Create a new node with the same value
		// (Next and Random will be set later)
		m[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// -------------------------
	// Pass 2: Assign Next and Random pointers
	// -------------------------
	current = head
	for current != nil {
		// m[current] is the cloned node
		// m[current.Next] gives the clone of the next node
		m[current].Next = m[current.Next]

		// m[current.Random] gives the clone of the random node
		m[current].Random = m[current.Random]

		current = current.Next
	}

	// Return the head of the cloned list
	// If head is nil, map lookup safely returns nil
	return m[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
