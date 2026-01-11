package neetcode150

// Problem: Reverse a singly linked list and return the new head.
//
// Pattern: Iterative Pointer Reversal (In-Place Linked List Manipulation)
//
// Intuition:
// A singly linked list can only be traversed forward.
// To reverse it, we must reverse the direction of each node's `Next` pointer.
// This is done one node at a time while carefully keeping track of:
//   - the previous node
//   - the current node
//   - the next node (to avoid losing the rest of the list)
//
// Key Insight:
// At every step, we "detach" the current node from the original list
// and point it backward to the already-reversed portion.
//
// Step-by-Step Idea:
// 1️⃣ Start with `prev = nil` (this will become the new tail).
// 2️⃣ Traverse the list using `current`.
// 3️⃣ Before changing pointers, save `current.Next`.
// 4️⃣ Reverse the pointer: `current.Next = prev`.
// 5️⃣ Move `prev` and `current` one step forward.
// 6️⃣ When traversal ends, `prev` is the new head.
//
// Time Complexity: O(n) — each node is visited once
// Space Complexity: O(1) — in-place, no extra memory used

func reverseList(head *ListNode) *ListNode {
	// `prev` will track the reversed portion of the list
	var prev *ListNode

	// `next` temporarily stores the next node to prevent losing the list
	var next *ListNode

	// `current` is used to traverse the list
	current := head

	for current != nil {
		// 🔒 Step 1: Save the next node
		next = current.Next

		// 🔁 Step 2: Reverse the pointer
		current.Next = prev

		// 👉 Step 3: Move `prev` forward
		prev = current

		// 👉 Step 4: Move `current` forward
		current = next
	}

	// `prev` now points to the new head of the reversed list
	return prev
}
