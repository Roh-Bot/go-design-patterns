package neetcode150

//Example 1:
//
//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
//Example 2:
//
//Input: s1 = "ab", s2 = "eidboaoo"
//Output: false

func checkInclusion(s1 string, s2 string) bool {
	var arrS1 [26]int
	for i := 0; i < len(s1); i++ {
		arrS1[s1[i]-'a']++
	}

	windowEnd := len(s1) - 1
	for windowStart := 0; windowEnd < len(s2); windowStart++ {
		var arrS2 [26]int
		//eidbaooo
		for current := windowStart; current <= windowEnd; current++ {
			arrS2[s2[current]-'a']++
		}

		if arrS2 == arrS1 {
			return true
		}
		windowEnd++
	}
	return false
}
