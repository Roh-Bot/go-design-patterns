package neetcode150

import "math"

// Problem: Search a Sorted 2D Matrix
// ---------------------------------
// You are given a matrix where:
//   1) Each row is sorted in ascending order.
//   2) The first element of each row is greater than the last of the previous.
//
// Example:
// [
//   [1, 3, 5, 7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 60]
// ]
// Searching for target = 16 → return true.
//
// Pattern Used: Binary Search Twice
// ---------------------------------
// Step 1 → Binary search **on rows** to determine which row MAY contain the target.
//          Because each row spans a disjoint numeric range.
// Step 2 → Once the row is found, do a normal binary search inside that row.
//
// Why this works:
// The matrix behaves like a “sorted list of sorted rows”.
// So we can first locate the only possible row (range check),
// and then apply standard binary search.
//
// Time Complexity: O(log m + log n) where m = rows, n = columns.
// Space Complexity: O(1)

func searchMatrix(matrix [][]int, target int) bool {
	// First binary search: find the row where target *may* exist.
	low := 0
	high := len(matrix) - 1
	rowToSearch := math.MaxInt // sentinel meaning "no row found"

loop:
	for low <= high {
		mid := low + (high-low)/2
		lastCol := len(matrix[mid]) - 1

		// Case 1: target is smaller than the first element → move upward
		if matrix[mid][0] > target && matrix[mid][lastCol] > target {
			high = mid - 1

			// Case 2: target is greater than the last element → move downward
		} else if matrix[mid][0] < target && matrix[mid][lastCol] < target {
			low = mid + 1

			// Case 3: current row’s range may contain the target → stop
		} else {
			rowToSearch = mid
			break loop
		}
	}

	// If no row was found, target cannot exist
	if rowToSearch == math.MaxInt {
		return false
	}

	// Second binary search: search inside the found row.
	low = 0
	high = len(matrix[rowToSearch]) - 1
	for low <= high {
		mid := low + (high-low)/2

		if matrix[rowToSearch][mid] < target {
			low = mid + 1
		} else if matrix[rowToSearch][mid] > target {
			high = mid - 1
		} else {
			return true // found it!
		}
	}

	return false // not found
}
