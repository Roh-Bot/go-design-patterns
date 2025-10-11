package neetcode150

//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Pattern: Sliding Window + HashSet
// Approach:
// - Use two pointers (left, i) to represent a sliding window.
// - Expand window by moving i, shrink window from left if duplicate is found.
// - Track characters in the window using a set (map[byte]struct{}).
// - Update the maximum window size as we go.
// Complexity: O(n) time, O(min(n, charset)) space.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left := 0                     // left pointer for the sliding window
	longest := 0                  // store the max length found
	m := make(map[uint8]struct{}) // hash set to track unique chars in the current window

	// Expand the window with right pointer `i`
	for i := 0; i < len(s); i++ {
		// If s[i] already exists in the set, shrink window from the left
		for {
			if _, ok := m[s[i]]; !ok {
				break // stop shrinking when s[i] is no longer in the set
			}
			delete(m, s[left]) // remove the leftmost character
			left++             // move the left pointer forward
		}

		// Add current character to the set
		m[s[i]] = struct{}{}

		// Update max length of substring without duplicates
		longest = max(longest, i-left+1)
	}

	return longest
}

// lengthOfLongestSubstring finds the length of the longest substring
// without repeating characters using the "jumping left" trick.
func lengthOfLongestSubstringLastSeenIndex(s string) int {
	// Map to store the last index where a character was seen
	lastSeen := make(map[byte]int)

	// Left pointer of sliding window
	left := 0

	// Maximum length found so far
	maxLen := 0

	// Iterate through the string with the right pointer
	for right := 0; right < len(s); right++ {
		char := s[right]

		// If char was seen and inside current window → move left
		if idx, found := lastSeen[char]; found && idx >= left {
			// Instead of left++ repeatedly,
			// we JUMP left to the position after the duplicate.
			left = idx + 1
		}

		// Update last seen index of current char
		lastSeen[char] = right

		// Update max length of substring found so far
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
