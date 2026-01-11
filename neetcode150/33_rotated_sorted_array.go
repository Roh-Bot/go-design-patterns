package neetcode150

// Problem: Search a target value in a rotated sorted array.
// Example: nums = [4,5,6,7,0,1,2], target = 0 → return 4
//
// Pattern: Modified Binary Search
//
// Idea:
// A rotated sorted array is two sorted halves glued together.
// Example rotation: [0,1,2,4,5,6,7] → [4,5,6,7,0,1,2]
//
// At any `mid` index, one side (left or right) is guaranteed to be sorted.
// We:
// 1️⃣ Identify which half is sorted
// 2️⃣ Check if the target lies inside the sorted half
// 3️⃣ If yes → binary search inside that half
//     If no → discard and search in the other half
//
// Why works?
// Even after rotation, one portion around `mid` is still sorted,
// enabling a deterministic binary-search-like reduction each step.
//
// Time Complexity: O(log n)
// Space Complexity: O(1)

func searchRotated(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		// Found target → return index
		if nums[mid] == target {
			return mid
		}

		// Check if LEFT half is sorted
		if nums[left] <= nums[mid] {
			// Is target inside left sorted portion?
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Otherwise RIGHT half is sorted
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	// Target not found
	return -1
}
