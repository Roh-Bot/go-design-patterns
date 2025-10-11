package neetcode150

// Example 2:
//
// Input: nums = [1,2,3,4]
//
// Output: false
//
// Explanation:
//
// All elements are distinct.
//
// Example 3:
//
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
//
// Output: true

// containsDuplicate checks if the given slice of integers contains any duplicates.
// Pattern Used: Hash Set Membership Check
// - We use a map (acting as a hash set) to keep track of numbers we've already seen.
// - If a number appears again while iterating, we immediately return true.
// - Otherwise, if no duplicates are found, return false.
func containsDuplicate(nums []int) bool {
	// Create a map to store numbers we've already encountered.
	// We only care about the existence of a number, not any value,
	// so we use `struct{}` as the value type (because it takes 0 bytes).
	m := make(map[int]struct{}, len(nums))

	// Loop through each number in the input slice
	for _, v := range nums {
		// Check if this number already exists in the map
		if _, ok := m[v]; ok {
			// If yes, we've found a duplicate → return true immediately
			return true
		}
		// Otherwise, store the number in the map
		m[v] = struct{}{}
	}

	// If we finish the loop without finding duplicates, return false
	return false
}
