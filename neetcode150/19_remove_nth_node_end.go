package neetcode150

// Problem: Remove Nth Node From End of List
// Given the head of a singly linked list, remove the nth node from the end
// and return the head of the modified list.
//
// Pattern:
// Two Pointers (Fast & Slow) with fixed distance
//
// Intuition:
// If we keep two pointers `first` and `second` such that the distance
// between them is `n` nodes, then when `first` reaches the end of the list,
// `second` will be exactly at the node we need to remove.
//
// Key Insight:
// Removing the nth node from the *end* is equivalent to removing the
// (length - n + 1)th node from the *start*, but we don’t know the length.
// Two pointers allow us to avoid computing length.
//
// Step-by-step Strategy:
//
// Step 1: Move `first` pointer `n` steps ahead
//   - This creates a gap of `n` nodes between `first` and `second`
//
// Step 2: Handle edge case
//   - If `first` becomes nil after moving `n` steps,
//     it means we need to remove the head node
//
// Step 3: Move both pointers together
//   - Keep track of `prev`, the node before `second`
//   - When `first` reaches the end, `second` is the target node
//
// Step 4: Remove the target node
//   - Bypass it using `prev.Next = second.Next`
//
// Time Complexity: O(n)
// Space Complexity: O(1)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Initialize two pointers at the head
	first := head
	second := head

	// Move `first` n steps ahead to create a gap of n nodes
	for range n {
		first = first.Next
	}

	// If `first` is nil after the shift,
	// it means the node to remove is the head itself
	if first == nil {
		return head.Next
	}

	// `prev` will track the node just before `second`
	prev := second

	// Move both pointers until `first` reaches the end
	for first != nil {
		prev = second
		second = second.Next
		first = first.Next
	}

	// `second` is now pointing to the nth node from the end
	// Remove it by skipping over it
	prev.Next = second.Next

	return head
}
