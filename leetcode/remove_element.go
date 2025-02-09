package leetcode

func RemoveElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
		i++
	}
	return count
}
