package easy

import "strings"

func LengthLastWord(s string) int {
	s1 := strings.Split(strings.TrimSpace(s), " ")
	return len(s1[len(s1)-1])
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			length++
			continue
		}
		break
	}
	return length
}
