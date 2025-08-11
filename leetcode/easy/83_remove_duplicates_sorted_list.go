package easy

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	// check if current is not null as
	// if current and its next val matches we
	// proceed to replace the current.next node by the current.next.next
	// and move the pointer with the same operation
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else { // If it doesn't match simply move the pointer by 1
			current = current.Next
		}
	}

	return head
}
