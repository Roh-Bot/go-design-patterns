package heaps

// MinHeap represents a binary heap where the smallest element is always at the root (index 0).
// It ensures that for every node i, arr[i] <= arr[2*i+1] and arr[i] <= arr[2*i+2].
type MinHeap struct {
	arr []int
}

// Insert adds a new element to the heap and performs "heapify up" (or "sift-up")
// to maintain the min-heap property.
// Time complexity: O(log n)
func (h *MinHeap) Insert(data int) {
	h.arr = append(h.arr, data) // add element at the end (maintains complete tree)

	currentIdx := len(h.arr) - 1

	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2

		// If parent is already smaller than or equal to child, property is valid.
		if h.arr[parentIdx] <= h.arr[currentIdx] {
			break
		}

		// Otherwise, swap with parent and move up.
		h.arr[currentIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[currentIdx]
		currentIdx = parentIdx
	}
}

// Pop removes and returns the smallest element (root of the heap).
// Steps:
// 1. Save the root (minimum element).
// 2. Move last element to root position and reduce heap size.
// 3. Restore heap property with heapifyDown.
// Time complexity: O(log n)
func (h *MinHeap) Pop() int {
	if len(h.arr) == 0 {
		return -1 // could return (int, error) in production
	}

	minNum := h.arr[0]

	// Move last element to root
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	// Restore min-heap property
	h.heapifyDown(0, len(h.arr)-1)

	return minNum
}

// heapifyDown restores the min-heap property starting from the given index.
// It swaps the parent with its smaller child until the heap property is satisfied.
// Time complexity: O(log n)
func (h *MinHeap) heapifyDown(currentParentIdx, lastIdx int) {
	for {
		leftChildIdx := (2 * currentParentIdx) + 1
		rightChildIdx := (2 * currentParentIdx) + 2

		// Assume the current node is the smallest
		smallestIdx := currentParentIdx

		// If left child exists and is smaller than the current smallest
		if leftChildIdx <= lastIdx && h.arr[leftChildIdx] < h.arr[smallestIdx] {
			smallestIdx = leftChildIdx
		}

		// If right child exists and is smaller than the current smallest
		if rightChildIdx <= lastIdx && h.arr[rightChildIdx] < h.arr[smallestIdx] {
			smallestIdx = rightChildIdx
		}

		// If no swaps needed, heap property satisfied
		if smallestIdx == currentParentIdx {
			break
		}

		// Swap and continue down
		h.arr[currentParentIdx], h.arr[smallestIdx] = h.arr[smallestIdx], h.arr[currentParentIdx]
		currentParentIdx = smallestIdx
	}
}

// Peek returns the smallest element (root of the heap) without removing it.
// Time complexity: O(1)
func (h *MinHeap) Peek() int {
	if len(h.arr) == 0 {
		return -1
	}
	return h.arr[0]
}

// Sort performs an in-place heap sort using the current heap array.
// After calling Sort(), the array will be sorted in descending order (since min-heap).
// Time complexity: O(n log n)
func (h *MinHeap) Sort() {
	n := len(h.arr) - 1

	for n > 0 {
		// Move smallest (root) to end
		h.arr[0], h.arr[n] = h.arr[n], h.arr[0]
		n--

		// Restore heap property in reduced heap
		h.heapifyDown(0, n)
	}
}

// Heapify converts an arbitrary slice into a valid min-heap in O(n) time.
// This is more efficient than inserting each element individually (O(n log n)).
func (h *MinHeap) Heapify(arr []int) []int {
	h.arr = arr
	n := len(arr)

	// Start from the last parent node and heapify down to root.
	for i := (n - 2) / 2; i >= 0; i-- {
		h.heapifyDown(i, n-1)
	}

	return h.arr
}
