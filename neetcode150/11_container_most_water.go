package neetcode150

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

// Pattern: Two-Pointer Technique (Greedy)
// Approach:
// - Start with the widest container (left=0, right=n-1).
// - Compute area = min(height[left], height[right]) * width.
// - Move the pointer at the shorter height inward, since the limiting factor is the shorter line.
// - Update the maximum area while moving inward.
// Complexity: O(n) time, O(1) space.
func maxArea(height []int) int {
	largestContainer := 0 // track the maximum water area found

	left := 0                // left pointer at the start
	right := len(height) - 1 // right pointer at the end

	// Move pointers toward each other
	for left < right {
		containerArea := 0

		// Area is limited by the shorter line between left and right
		if height[left] < height[right] {
			// Width = (right - left), Height = height[left]
			containerArea = height[left] * (right - left)
			left++ // move left pointer inward (hoping for a taller line)
		} else {
			// Width = (right - left), Height = height[right]
			containerArea = height[right] * (right - left)
			right-- // move right pointer inward
		}

		// Update maximum area
		largestContainer = max(largestContainer, containerArea)
	}

	return largestContainer
}
