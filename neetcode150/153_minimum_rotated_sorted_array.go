package neetcode150

// Problem: Find the minimum element in a rotated sorted array.
// Example:
//   nums = [3,4,5,1,2] → 1
//   nums = [11,13,15,17] → 11
//
// Pattern: Modified Binary Search (Rotated Array)
//
// Intuition:
// A rotated sorted array consists of two sorted halves.
// The minimum element is the "pivot" — the only element
// where the sorted order breaks.
//
// Key Observations:
// 1️⃣ If nums[left] < nums[right], the subarray is already sorted.
//     → The minimum is simply nums[left].
// 2️⃣ Otherwise, the minimum lies somewhere inside.
// 3️⃣ At every step, one half is guaranteed to be sorted.
//     - If left..mid is sorted → pivot must be on the right.
//     - Else → pivot is on the left.
//
// We shrink the search space just like binary search.
//
// Time Complexity: O(log n)
// Space Complexity: O(1)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	// Initialize with first element
	minimum := nums[0]

	for left <= right {

		// Case 1: Current window is already sorted
		// The leftmost element is the minimum
		if nums[left] < nums[right] {
			minimum = min(minimum, nums[left])
			break
		}

		// Calculate mid
		mid := left + (right-left)/2

		// Update minimum with mid element
		minimum = min(minimum, nums[mid])

		// Case 2: Left half is sorted
		// Pivot must be in the right half
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			// Case 3: Right half is sorted
			// Pivot must be in the left half
			right = mid - 1
		}
	}

	return minimum
}
