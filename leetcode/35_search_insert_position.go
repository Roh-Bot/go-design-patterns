package leetcode

func SearchInsertPosition(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		default:
			return mid
		}
	}
	return low
}
