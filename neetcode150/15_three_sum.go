package neetcode150

import "slices"

//Example 1:
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Explanation:
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
//The distinct triplets are [-1,0,1] and [-1,-1,2].
//Notice that the order of the output and the order of the triplets does not matter.
//Example 2:
//
//Input: nums = [0,1,1]
//Output: []
//Explanation: The only possible triplet does not sum up to 0.
//Example 3:
//
//Input: nums = [0,0,0]
//Output: [[0,0,0]]
//Explanation: The only possible triplet sums up to 0.

// threeSum finds all unique triplets in the array which give the sum of zero.
// Approach:
//  1. Sort the array.
//  2. Fix one number (nums[i]) and use two pointers (left, right)
//     to find pairs that sum up with nums[i] to zero.
//  3. Skip duplicates to avoid repeating the same triplets.
func threeSum(nums []int) [][]int {
	// Step 1: sort numbers to make two-pointer search possible
	slices.Sort(nums)

	var result [][]int

	// Step 2: iterate through each number as the fixed element
	for i := 0; i < len(nums); i++ {
		// Skip duplicates for the fixed element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		current := nums[i]
		left := i + 1
		right := len(nums) - 1

		// Step 3: use two-pointer approach to find complementary pairs
		for left < right {
			sum := current + nums[left] + nums[right]

			switch {
			case sum > 0:
				right-- // too large → move right pointer left
			case sum < 0:
				left++ // too small → move left pointer right
			default:
				// Found a valid triplet
				result = append(result, []int{current, nums[left], nums[right]})
				left++

				// Skip duplicate numbers on the left
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return result
}
