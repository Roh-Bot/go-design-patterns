package easy

// ["flight","flower","flow","n"]
// ["dog","racecar","car"]
func longestCommonPrefix(str []string) string {
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for len(prefix) >= len(str[i]) && str[i][0:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func LongestCommonPrefix(str []string) string {
	return longestCommonPrefix(str)
}
