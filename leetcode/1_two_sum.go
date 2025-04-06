package leetcode

// nums=[2,7,3,11], target=9
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, v := range nums {
		num := target - v

		if numsIdx, ok := numsMap[num]; ok {
			return []int{numsIdx, i}
		}

		nums[v] = i
	}
	return nil
}
