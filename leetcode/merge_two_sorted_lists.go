package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	current := newList
	for list1 != nil || list2 != nil {
		if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	return newList.Next
}
