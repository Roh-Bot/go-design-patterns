package leetcode

func IndexFirstOccurrenceInString(haystack, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if len(needle)+i > len(haystack) {
			return -1
		} else if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
