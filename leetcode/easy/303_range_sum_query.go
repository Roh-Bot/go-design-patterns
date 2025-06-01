package easy

// INPUT
// [1,3,5,7,8,12]
// [1,4,9,16,24,36]
//sumRange(1, 3); // return 3 + 5 + 7 = 15

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sumArr := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sumArr[i] = nums[i] + sum
		sum += nums[i]
	}
	return NumArray{nums: sumArr}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.nums[right]
	}
	return n.nums[right] - n.nums[left-1]
}
