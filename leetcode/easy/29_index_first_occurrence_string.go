package easy

func IndexFirstOccurrenceInString(haystack, needle string) int {
	window := len(needle)
	for i := 0; window <= len(haystack); i++ {
		if haystack[i:window] == needle {
			return i
		}
		window++
	}
	return -1
}
