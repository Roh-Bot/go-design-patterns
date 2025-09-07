package neetcode150

//Example 1:
//
//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
//Example 2:
//
//Input: numbers = [2,3,4], target = 6
//Output: [1,3]
//Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
//Example 3:
//
//Input: numbers = [-1,0], target = -1
//Output: [1,2]
//Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

// Pattern: Two-Pointer Technique on Sorted Array
// Approach:
// - Place one pointer at the start, another at the end.
// - If sum is too large → move right pointer left.
// - If sum is too small → move left pointer right.
// - Stop when the target sum is found.
// Complexity: O(n) time, O(1) space.
func twoSumII(numbers []int, target int) []int {
	// Step 1: Initialize two pointers:
	// - i starts at the beginning of the array
	// - j starts at the end of the array
	i := 0
	j := len(numbers) - 1

	// Step 2: While pointers don’t cross each other
	for i < len(numbers)-1 && j >= 0 {
		sum := numbers[i] + numbers[j] // sum of the two numbers

		switch {
		case sum > target:
			// If sum is too big, move the right pointer left
			// (because numbers[j] is too large)
			j--
		case sum < target:
			// If sum is too small, move the left pointer right
			// (because numbers[i] is too small)
			i++
		default:
			// Found the pair!
			// Problem statement requires 1-based indices
			return []int{i + 1, j + 1}
		}
	}

	// If no solution found (though problem guarantees one), return nil
	return nil
}
