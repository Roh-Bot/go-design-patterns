package heaps

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(data int) {
	h.arr = append(h.arr, data)

	currentIdx := len(h.arr) - 1

	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2

		if h.arr[parentIdx] < h.arr[currentIdx] {
			break
		}
		h.arr[currentIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[currentIdx]
		currentIdx = parentIdx
	}
}

func (h *MinHeap) Pop() int {
	minNum := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0, len(h.arr)-1)

	return minNum
}

func (h *MinHeap) heapifyDown(currentParentIdx, lastIdx int) {
	for {
		leftChildIdx := (currentParentIdx * 2) + 1
		rightChildIdx := (currentParentIdx * 2) + 2

		// assuming minimum index in current parent
		// where the parent can be a leaf node and
		// where the parent is already the minimum
		smallestIdx := currentParentIdx

		if leftChildIdx <= lastIdx && h.arr[leftChildIdx] < h.arr[currentParentIdx] {
			smallestIdx = leftChildIdx
		}
		if rightChildIdx <= lastIdx && h.arr[rightChildIdx] < h.arr[currentParentIdx] {
			smallestIdx = rightChildIdx
		}

		if smallestIdx == currentParentIdx {
			break
		}

		h.arr[smallestIdx], h.arr[currentParentIdx] = h.arr[currentParentIdx], h.arr[smallestIdx]
		currentParentIdx = smallestIdx
	}
}

func (h *MinHeap) Peek() int {
	return h.arr[len(h.arr)-1]
}

func (h *MinHeap) Sort() {
	n := len(h.arr) - 1
	for n > 0 {
		h.arr[0], h.arr[n] = h.arr[n], h.arr[0]
		n--
		h.heapifyDown(0, n)
	}
}

func (h *MinHeap) Heapify(arr []int) []int {
	h.arr = arr
	n := len(arr)

	for i := (n - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i, n-1)
	}
	return arr
}
