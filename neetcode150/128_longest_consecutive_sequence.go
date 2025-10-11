package neetcode150

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
// Example 3:
//
// Input: nums = [1,0,1,2]
// Output: 3

// Pattern: HashSet for O(1) lookups
// Approach:
// - Insert all numbers into a set.
// - Only start counting when a number is the beginning of a sequence (no v-1).
// - Extend streak by checking v+1, v+2… until sequence breaks.
// Complexity: O(n) time, O(n) space.
func longestConsecutive(nums []int) int {
	// Step 1: Put all numbers into a hash set (map with empty struct)
	// Why? O(1) lookup time for checking if a number exists
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}

	longest := 0 // will track the longest streak length

	// Step 2: Loop through each number in the set
	for v := range m {
		// If v-1 does not exist, that means `v` is the START of a sequence
		// Example: if v=5 and 4 is not in the set → 5 must start a sequence
		if _, hasPrev := m[v-1]; hasPrev {
			continue
		}

		curr := v   // current number we are exploring
		streak := 1 // current streak length

		// Step 3: Keep checking if consecutive numbers exist (v+1, v+2…)
		for {
			if _, isNextElementFound := m[curr+1]; !isNextElementFound {
				break // stop if next element doesn’t exist
			}
			curr++   // move to next number
			streak++ // increase streak count
		}

		// Step 4: Update longest streak if current streak is bigger
		longest = max(streak, longest)
	}
	return longest
}
