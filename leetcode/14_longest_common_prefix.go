package leetcode

import (
	"strings"
)

// ["flight","flower","flow","nikig]
// ["dog","racecar","car"]
func longestCommonPrefix(str []string) string {
	prefix := str[0]
	for i := 1; i < len(str); i++ {
		for !strings.HasPrefix(str[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func LongestCommonPrefix(str []string) string {
	return longestCommonPrefix(str)
}
