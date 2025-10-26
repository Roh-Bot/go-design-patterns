package neetcode150

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
// Example 1:
//
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]

// topKFrequent returns the k most frequent elements from nums.
// Pattern Used: Hash Map Frequency Count + Bucket Sort
// - Step 1: Count frequency of each number using a hash map.
// - Step 2: Place each number into a "bucket" where the index = frequency.
// - Step 3: Starting from the highest frequency bucket, collect numbers until we have k.
func topKFrequent(nums []int, k int) []int {
	// Step 1: Count frequency of each number
	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v]++ // increment frequency for number v
	}

	// Step 2: Create buckets where index = frequency of numbers
	// Since frequency can't exceed len(nums), we allocate len(nums)+1 buckets
	buckets := make([][]int, len(nums)+1)

	// Fill the buckets with numbers based on their frequency
	for num, count := range numsMap {
		// Place 'num' in the bucket corresponding to its frequency
		buckets[count] = append(buckets[count], num)
	}

	// Step 3: Collect the top K frequent elements
	var result []int
	currCount := 0

	// Start from the highest frequency down to lowest
loop:
	for i := len(buckets) - 1; i >= 0; i-- {
		// Skip empty buckets
		if buckets[i] == nil {
			continue
		}
		// For each number in this bucket
		for _, num := range buckets[i] {
			if currCount == k {
				// Stop once we've collected k numbers
				break loop
			}
			result = append(result, num)
			currCount++
		}
	}

	return result
}
