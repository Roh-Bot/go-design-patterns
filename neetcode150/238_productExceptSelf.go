package neetcode150

//Prefix: nums = [1,1,2,6]
//Prefix: nums = [24,12,4,1]

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// productExceptSelf returns an array where each element at index i
// is the product of all elements in nums except nums[i].
//
// Pattern Used: Prefix & Postfix Multiplication
// - Step 1: Build a prefix product for each index (product of all numbers before it).
// - Step 2: Build a postfix product for each index (product of all numbers after it).
// - Step 3: Multiply prefix and postfix for each index to get the final result.
// This avoids division and works in O(n) time and O(1) extra space.
func productExceptSelf(nums []int) []int {
	// Create the result array to hold the final answers
	result := make([]int, len(nums))

	// Step 1: Build prefix products
	// prefix holds the product of all numbers before index i
	prefix := 1
	result[0] = prefix // no numbers before index 0, so prefix = 1

	// Loop left-to-right, filling result[i] with product of numbers before i
	for i := 1; i < len(nums); i++ {
		// result[i] = prefix product up to (i-1)
		result[i] = prefix * nums[i-1]
		// update prefix to include nums[i-1]
		prefix = prefix * nums[i-1]
	}

	// Step 2: Multiply postfix products
	// postfix holds the product of all numbers after index i
	postfix := 1
	// Loop right-to-left
	for i := len(nums) - 1; i >= 0; i-- {
		// Multiply current value (prefix product) by postfix product
		result[i] = result[i] * postfix
		// update postfix to include nums[i]
		postfix = postfix * nums[i]
	}

	return result
}
