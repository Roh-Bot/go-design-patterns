package neetcode150

// maxSlidingWindow returns an array where each element is the maximum value
// within a sliding window of size `k` as it moves across `nums`.
//
// The function uses a deque (double-ended queue) to keep track of indices
// of elements in decreasing order, ensuring the largest element of the
// current window is always at the front.
//
// Time Complexity: O(n)
// - Each element is pushed and popped from the deque at most once.
// Space Complexity: O(k)
// - The deque holds at most k indices.
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	var result []int
	d := new(Dequeue) // custom deque to store indices

	for i := 0; i < len(nums); i++ {
		// Step 1: Remove indices that are out of the current window range.
		// (i - k) represents the leftmost index that just slid out.
		if d.Length() > 0 && d.PeekStart() <= i-k {
			d.PopStart()
		}

		// Step 2: Maintain decreasing order of elements in deque.
		// Remove all indices whose corresponding values are smaller
		// than the current number — they can never be window maximums.
		for d.Length() > 0 && nums[i] >= nums[d.PeekEnd()] {
			d.PopEnd()
		}

		// Step 3: Add current index to the deque (to the right end).
		d.PushEnd(i)

		// Step 4: Once we've reached at least window size `k`,
		// record the value at the front (which is the maximum).
		if i >= k-1 {
			result = append(result, nums[d.PeekStart()])
		}
	}

	return result
}
