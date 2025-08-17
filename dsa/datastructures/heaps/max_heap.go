package heaps

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(data int) {
	h.arr = append(h.arr, data)

	currentIdx := len(h.arr) - 1
	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2

		if h.arr[parentIdx] > h.arr[currentIdx] {
			break
		}

		h.arr[parentIdx], h.arr[currentIdx] = h.arr[currentIdx], h.arr[parentIdx]
		currentIdx = parentIdx
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}

	maxNumber := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0, len(h.arr)-1)

	return maxNumber
}

func (h *MaxHeap) heapifyDown(currentParentIdx int, lastIdx int) {
	for {
		leftChildIdx := (2 * currentParentIdx) + 1
		rightChildIdx := (2 * currentParentIdx) + 2

		// assuming parent is the greatest
		// where parent can be a leaf node and the case
		// where the parent is already the largest
		largestIdx := currentParentIdx

		if leftChildIdx <= lastIdx && h.arr[leftChildIdx] > h.arr[currentParentIdx] {
			largestIdx = leftChildIdx
		}
		if rightChildIdx <= lastIdx && h.arr[rightChildIdx] > h.arr[currentParentIdx] {
			largestIdx = rightChildIdx
		}

		if largestIdx == currentParentIdx {
			break
		}

		h.arr[currentParentIdx], h.arr[largestIdx] = h.arr[largestIdx], h.arr[currentParentIdx]
		currentParentIdx = largestIdx
	}
}

func (h *MaxHeap) Peek() int {
	return h.arr[len(h.arr)-1]
}

func (h *MaxHeap) Sort() {
	// Hypothetical last index
	currentLastIdx := len(h.arr) - 1

	for currentLastIdx > 0 {
		h.arr[0], h.arr[currentLastIdx] = h.arr[currentLastIdx], h.arr[0]
		currentLastIdx--
		h.heapifyDown(0, currentLastIdx)
	}
}

func (h *MaxHeap) Heapify(arr []int) []int {
	h.arr = arr
	n := len(arr)

	for i := (n - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i, n-1)
	}
	return arr
}
