package easy

// Input: s = "anagram", t = "nagaram"
//
// Output: true
//
// Example 2:
//
// Input: s = "rat", t = "car"
//
// Output: false
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var stats [26]int

	for i := 0; i < len(s); i++ {
		stats[s[i]-'a']++
		stats[t[i]-'a']--
	}

	for _, count := range stats {
		if count != 0 {
			return false
		}
	}

	return true
}
