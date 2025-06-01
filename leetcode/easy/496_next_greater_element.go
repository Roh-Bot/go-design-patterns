package easy

type MonotonicStack []int

func (m *MonotonicStack) push(data int) {
	for !m.IsEmpty() && data > (*m)[len(*m)-1] {
		m.pop()
	}
	*m = append(*m, data)
}

func (m *MonotonicStack) pop() {
	*m = (*m)[:len(*m)-1]
}

func (m *MonotonicStack) IsEmpty() bool {
	return len(*m) == 0
}

func (m *MonotonicStack) Peek() int {
	return (*m)[len(*m)-1]
}

// Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
// Output: [-1,3,-1]
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = i
	}

	monoStack := make(MonotonicStack, 0)

loop:
	for i := len(nums2) - 1; i >= 0; i-- {
		for !monoStack.IsEmpty() {
			if nums2[i] >= monoStack.Peek() {
				monoStack.pop()
				continue
			}

			if v, ok := nums1Map[nums2[i]]; ok {
				nums1[v] = monoStack.Peek()
			}

			monoStack.push(nums2[i])
			continue loop
		}

		if v, ok := nums1Map[nums2[i]]; ok {
			nums1[v] = -1
		}
		monoStack.push(nums2[i])
	}
	return nums1
}

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement(nums1, nums2)
}
