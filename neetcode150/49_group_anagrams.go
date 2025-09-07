package neetcode150

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Pattern: Hash Map Grouping (Character Frequency Signature)
// Uses a fixed-size array of 26 to represent character counts in each string.
// This frequency signature acts as a unique key for grouping anagrams.
// The map stores signature -> index in result, so each group of anagrams
// can be collected efficiently.
// Time Complexity: O(n * k), where n is number of strings and k is average string length.
// Space Complexity: O(n * k).

func groupAnagrams(strs []string) [][]string {
	// If there's only one string, just return it as a single group.
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	// This will hold the final grouped anagrams.
	// Example: [["eat","tea","ate"], ["bat"], ["tan","nat"]]
	var result [][]string

	// Map to keep track of groups we've already created.
	// Key   = [26]int (frequency of letters a-z in a string)
	// Value = index of the group inside 'result'
	m := make(map[[26]int]int)

	// Loop through each string in the input list
	for _, str := range strs {
		// Create a "frequency signature" for the current string.
		// stats[i] will store how many times the letter 'a'+i appears in 'str'.
		var stats [26]int

		// Loop through each character in the string
		for _, v := range str {
			// Subtracting 'a' from the character gives us its position in the alphabet (0–25).
			// Example: 'a'-'a' = 0, 'b'-'a' = 1, ..., 'z'-'a' = 25
			stats[v-'a']++
		}

		// Now 'stats' is the signature of the string.
		// Example: "eat" -> stats = [1,0,0,...,1,...,1,...] (with counts for 'a','e','t').

		// Check if this signature already exists in the map
		if index, ok := m[stats]; ok {
			// If yes, append the string to the existing group in 'result'
			result[index] = append(result[index], str)
			continue
		}

		// Otherwise, this is the first time we see this signature.
		// So we create a new group with this string.
		result = append(result, []string{str})

		// Store the index of this new group in the map for future lookups.
		m[stats] = len(result) - 1
	}

	// Return all groups of anagrams.
	return result
}
