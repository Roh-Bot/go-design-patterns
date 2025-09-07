package neetcode150

// Input: nums=[2,7,3,11], target=9
// Pattern: Hash Map Lookup (Complement Map Technique)
// This uses a map to store previously seen numbers and their indices.
// For each number, we check if its complement (target - current number)
// is already in the map. If yes, return the indices; otherwise, store the number.
// Time Complexity: O(n), Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, v := range nums {
		num := target - v

		if numsIdx, ok := numsMap[num]; ok {
			return []int{numsIdx, i}
		}

		numsMap[v] = i
	}
	return nil
}
