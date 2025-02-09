package leetcode

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (n *NumArray) SumRange(left int, right int) int {
	n.processArray()
	if left == 0 {
		return n.nums[right]
	}
	return n.nums[right] - n.nums[left-1]
}

func (n *NumArray) processArray() {
	for i := 1; i < len(n.nums); i++ {
		n.nums[i] = n.nums[i-1] + n.nums[i]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
