package easy

// INPUT
// [1,3,5,7,8,12]
// [1,4,9,16,24,36]
//sumRange(1, 3); // return 3 + 5 + 7 = 15

// PATTERN USED: PREFIX SUM
// Preprocess the array to create a new
// array where each element at index i
// represents the sum of the array from the start up to i
// then simply use the formula P[j] - P[i-1] to calculate the
// sum between two indices

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	sumArr := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sumArr[i] += nums[i]
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
