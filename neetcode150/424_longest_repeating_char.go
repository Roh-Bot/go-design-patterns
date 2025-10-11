package neetcode150

//Input: s = "ABAB", k = 2
//Output: 4
//Explanation: Replace the two 'A's with two 'B's or vice versa.
//Example 2:
//
//Input: s = "AABABBA", k = 1
//Output: 4
//Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
//The substring "BBBB" has the longest repeating letters, which is 4.
//There may exists other ways to achieve this answer too.

// characterReplacement finds the length of the longest substring
// where you can replace at most k characters so that all characters
// in the substring are the same.
// Optimized: uses a fixed-size [26]int frequency array instead of a map
// for faster lookups (since only uppercase A–Z are involved).
func characterReplacement(s string, k int) int {
	var arr [26]int // frequency array for A–Z characters

	maxOccurrence := 0 // most frequent char count in the window
	left := 0
	maxLen := 1

	for right := 0; right < len(s); right++ {
		// increment frequency of current char
		arr[s[right]-'A']++
		currentOccurrence := arr[s[right]-'A']

		// update most frequent char count if needed
		maxOccurrence = max(maxOccurrence, currentOccurrence)

		// shrink window if replacements needed > k
		for right-left+1-maxOccurrence > k {
			arr[s[left]-'A']--
			left++
		}

		// update max length found so far
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// characterReplacement finds the length of the longest substring
// where you can replace at most k characters so that all characters
// in the substring are the same.
// Optimized: we maintain mostOccurred directly, instead of rescanning the map each time.
func characterReplacementNoRescan(s string, k int) int {
	m := make(map[byte]int) // frequency map for chars in window

	left := 0
	maxLen := 0
	mostOccurred := 0 // count of the most frequent char in the current window

	for right := 0; right < len(s); right++ {
		char := s[right]
		m[char]++

		// update max frequency if this char became more frequent
		if m[char] > mostOccurred {
			mostOccurred = m[char]
		}

		// current window size
		windowLength := right - left + 1

		// if replacements needed > k, shrink window
		if windowLength-mostOccurred > k {
			m[s[left]]--
			left++
			windowLength = right - left + 1
		}

		// update answer
		if windowLength > maxLen {
			maxLen = windowLength
		}
	}

	return maxLen
}
