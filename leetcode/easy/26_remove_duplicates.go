package easy

// Input: [1,1,2,2,3,4,5]
func RemoveDuplicatesFromSortedArray(nums []int) int {
	nextInsertionIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextInsertionIndex] = nums[i]
			nextInsertionIndex++
		}
	}

	return nextInsertionIndex
}
