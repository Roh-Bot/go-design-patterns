package heaps

// MaxHeap holds a slice of ints where the largest element is always at index 0.
type MaxHeap struct {
	arr []int
}

// Insert adds a new element to the heap and then performs "heapify up"
// (sometimes called sift-up or bubble-up) to preserve the max-heap property.
// Complexity: O(log n) in the height of the heap.
func (h *MaxHeap) Insert(data int) {
	// Append new element at the end (a leaf).
	h.arr = append(h.arr, data)

	// Start at the newly inserted node and move it up until parent >= current.
	currentIdx := len(h.arr) - 1
	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2

		// If parent is already greater or equal, heap property satisfied.
		if h.arr[parentIdx] > h.arr[currentIdx] {
			break
		}

		// Otherwise swap with parent and continue upward.
		h.arr[parentIdx], h.arr[currentIdx] = h.arr[currentIdx], h.arr[parentIdx]
		currentIdx = parentIdx
	}
}

// Pop removes and returns the maximum element (root of the heap).
// Typical steps:
// 1. Save root value to return.
// 2. Move last element to root position, shrink slice.
// 3. heapifyDown from root to restore heap property.
// Complexity: O(log n).
// Returns -1 if heap is empty (your original code used -1, keep that but consider returning (int, error) in production).
func (h *MaxHeap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}

	// root (maximum) value
	maxNumber := h.arr[0]

	// move last element to root and shrink the underlying slice
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	// restore heap property from root downwards
	h.heapifyDown(0, len(h.arr)-1)

	return maxNumber
}

// heapifyDown restores the max-heap property starting at currentParentIdx
// and working down toward the leaves. lastIdx is the last valid index in h.arr.
// Complexity: O(log n).
func (h *MaxHeap) heapifyDown(currentParentIdx int, lastIdx int) {
	for {
		leftChildIdx := (2 * currentParentIdx) + 1
		rightChildIdx := (2 * currentParentIdx) + 2

		// assume current parent is largest; try to find a larger child
		largestIdx := currentParentIdx

		// if left exists and is greater than current largest, update largest
		if leftChildIdx <= lastIdx && h.arr[leftChildIdx] > h.arr[largestIdx] {
			largestIdx = leftChildIdx
		}
		// IMPORTANT: compare right child with h.arr[largestIdx] (not parent)
		if rightChildIdx <= lastIdx && h.arr[rightChildIdx] > h.arr[largestIdx] {
			largestIdx = rightChildIdx
		}

		// if no child was larger than parent, heap property satisfied
		if largestIdx == currentParentIdx {
			break
		}

		// swap parent with the larger child and continue
		h.arr[currentParentIdx], h.arr[largestIdx] = h.arr[largestIdx], h.arr[currentParentIdx]
		currentParentIdx = largestIdx
	}
}

// Peek returns the maximum element without removing it.
// Returns -1 if the heap is empty to avoid panics.
func (h *MaxHeap) Peek() int {
	if len(h.arr) == 0 {
		return -1
	}
	return h.arr[0]
}

// Sort performs an in-place heap sort using the heap array.
// After calling Sort(), h.arr will be sorted in ascending order.
// Complexity: O(n log n).
func (h *MaxHeap) Sort() {
	// last index available for swapping during sort
	currentLastIdx := len(h.arr) - 1

	// repeatedly move current max (root) to the end, reduce the heap size,
	// and heapify down the root for the reduced heap.
	for currentLastIdx > 0 {
		h.arr[0], h.arr[currentLastIdx] = h.arr[currentLastIdx], h.arr[0]
		currentLastIdx--
		h.heapifyDown(0, currentLastIdx)
	}
}

// Heapify takes an arbitrary slice and transforms it into a max-heap in O(n).
// It does so by calling heapifyDown from the last non-leaf node down to the root.
func (h *MaxHeap) Heapify(arr []int) []int {
	h.arr = arr
	n := len(arr)

	// last parent is (n-2)/2; walk backwards to root
	for i := (n - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i, n-1)
	}
	return h.arr
}
