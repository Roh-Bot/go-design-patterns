package easy

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var arr []uint8
	for i := 0; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			arr = append(arr, s[i])
		}
	}
	palindrome := strings.ToLower(string(arr))
	fmt.Println(palindrome)

	i := 0
	j := len(palindrome) - 1
	for i < len(palindrome) && j >= 0 {
		fmt.Printf("i: %s, j: %s", string(palindrome[i]), string(palindrome[j]))
		if palindrome[i] != palindrome[j] {
			return false
		}
		i++
		j--
	}
	return true
}
