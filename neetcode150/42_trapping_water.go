package neetcode150

// Input: height = [4,2,0,3,2,5]
// // Output: 6

// trapArrOn calculates trapped rain water using precomputed max arrays.
// Approach:
// 1. For each position, precompute the maximum height to the left and right.
// 2. Water trapped at position i = min(maxLeft[i], maxRight[i]) - height[i].
// Complexity: O(n) time, O(n) space.
func trapArrOn(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	maxLeftArr := make([]int, n) // max height to the left of each bar
	maxLeftArr[0] = height[0]
	maxLeft := height[0]
	for i := 1; i < n; i++ {
		maxLeft = max(height[i], maxLeft)
		maxLeftArr[i] = maxLeft
	}

	maxRightArr := make([]int, n) // max height to the right of each bar
	maxRightArr[n-1] = height[n-1]
	maxRight := height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight = max(height[i], maxRight)
		maxRightArr[i] = maxRight
	}

	totalTraps := 0
	for i := 0; i < n; i++ {
		totalTraps += min(maxLeftArr[i], maxRightArr[i]) - height[i]
	}
	return totalTraps
}

// trapTwoPointersO1 calculates trapped rain water using two pointers in O(1) space.
// Approach:
// 1. Keep two pointers at the ends and track max heights from left and right.
// 2. Always move the pointer with the smaller max height.
// 3. Water trapped at each step = maxHeight - current height.
// Complexity: O(n) time, O(1) space.
func trapTwoPointersO1(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	totalTraps := 0

	for left < right {
		if leftMax <= rightMax {
			left++
			totalTraps += max(0, leftMax-height[left])
			leftMax = max(leftMax, height[left])
		} else {
			right--
			totalTraps += max(0, rightMax-height[right])
			rightMax = max(rightMax, height[right])
		}
	}
	return totalTraps
}
