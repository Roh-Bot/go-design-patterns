package neetcode150

// Input: s = "anagram", t = "nagaram"
//
// Output: true
//
// Example 2:
//
// Input: s = "rat", t = "car"
//
// Output: false

// isAnagram checks if two strings s and t are anagrams of each other.
// Pattern Used: Hash Map / Frequency Array Comparison
//   - Instead of sorting both strings (which costs O(n log n)),
//     we count the frequency of each character using a fixed-size array of length 26.
//   - We increment counts for characters in s and decrement for characters in t.
//   - If the two strings are anagrams, the counts for all letters must balance to 0.
func isAnagram(s string, t string) bool {
	// Quick check: if the lengths differ, they can't be anagrams
	if len(s) != len(t) {
		return false
	}

	// stats[i] represents the net count of the (i+'a')th character.
	// Positive value = appears more times in s
	// Negative value = appears more times in t
	var stats [26]int

	// Process both strings in a single loop
	for i := 0; i < len(s); i++ {
		// Increment for character from s
		stats[s[i]-'a']++
		// Decrement for character from t
		stats[t[i]-'a']--
	}

	// After processing both strings:
	// - If they are true anagrams, all counts should be zero
	// - If not, some counts will be non-zero
	for _, count := range stats {
		if count != 0 {
			return false
		}
	}

	// If we didn't find any imbalance, the strings are anagrams
	return true
}
