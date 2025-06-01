package easy

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for {
		if current == nil || current.Next == nil {
			break
		}

		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
