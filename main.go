package main

func main() {
	nextGreaterElement([]int{4, 1, 2, 6}, []int{1, 3, 4, 2, 6, 7})
}

type MonotonicStack[T comparable] []T

func (m *MonotonicStack[T]) Peek() T {
	return (*m)[len(*m)-1]
}

func (m *MonotonicStack[T]) Push(data T) {
	*m = append(*m, data)
}

func (m *MonotonicStack[T]) Pop() {
	*m = (*m)[:len(*m)-1]
}

func (m *MonotonicStack[T]) IsEmpty() bool {
	return len(*m) == 0
}

// Input: nums1 = [4,1,2,6], nums2 = [1,3,4,2,6,7]
// Output: [-1,3,-1]
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)

	for i, v := range nums1 {
		nums1Map[v] = i
	}

	monoStack := make(MonotonicStack[int], 0)

loop:
	for i := len(nums2) - 1; i >= 0; i-- {
		for !monoStack.IsEmpty() {
			if nums2[i] > monoStack.Peek() {
				monoStack.Pop()

			} else if idx, ok := nums1Map[nums2[i]]; ok {
				nums1[idx] = monoStack.Peek()
			}
			monoStack.Push(nums2[i])
			continue loop
		}

		if idx, ok := nums1Map[nums2[i]]; ok {
			nums1[idx] = -1
		}
		monoStack.Push(nums2[i])
	}
	return nums2
}
