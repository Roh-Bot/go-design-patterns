package easy

func RemoveElement(nums []int, val int) int {
	insertionIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[insertionIdx] = nums[i]
			insertionIdx++
		}
	}
	return insertionIdx
}
