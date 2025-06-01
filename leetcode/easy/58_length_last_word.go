package easy

import "strings"

func LengthLastWord(s string) int {
	s1 := strings.Split(strings.TrimSpace(s), " ")
	return len(s1[len(s1)-1])
}
