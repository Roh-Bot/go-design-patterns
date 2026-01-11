package neetcode150

// Problem: Merge two sorted singly linked lists into one sorted list.
//
// Pattern: Dummy Node + Two-Pointer Merge
//
// Intuition:
// This is identical to the "merge" step of Merge Sort.
// Since both lists are already sorted, we can repeatedly take
// the smaller head node and append it to the result list.
//
// Key Insight:
// Using a dummy (sentinel) node eliminates edge cases like:
//   - initializing the head of the merged list
//   - handling the first insertion separately
//
// We maintain:
//   - `current` → tail of the merged list
//   - `list1` and `list2` → pointers traversing the input lists
//
// At every step, we:
//   - Compare current nodes of both lists
//   - Append the smaller node
//   - Advance the pointer of the list we used
//
// Time Complexity: O(n + m) — each node is visited once
// Space Complexity: O(1) — no new nodes created (reuses existing ones)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Dummy node acts as a stable starting point for the merged list
	dummy := &ListNode{}
	current := dummy

	// Traverse both lists until one is exhausted
	for list1 != nil && list2 != nil {
		// Always attach the smaller node
		if list1.Val <= list2.Val {
			current.Next = list1 // link list1 node
			list1 = list1.Next   // move list1 forward
		} else {
			current.Next = list2 // link list2 node
			list2 = list2.Next   // move list2 forward
		}
		current = current.Next // move merged list forward
	}

	// Attach the remaining nodes (only one list can be non-nil)
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// The merged list starts after the dummy node
	return dummy.Next
}
