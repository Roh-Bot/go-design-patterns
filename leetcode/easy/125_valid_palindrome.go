package easy

import (
	"fmt"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into
// lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
func isPalindrome(s string) bool {
	// hold only letters in arr
	// comparisons done with string represent their ASCII DEC value
	var arr []uint8
	for i := 0; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			arr = append(arr, s[i])
		}
	}
	palindrome := strings.ToLower(string(arr))
	fmt.Println(palindrome)

	// Take two pointers, one from start and other from end
	// return false when they mismatch
	i := 0
	j := len(palindrome) - 1
	for i < len(palindrome) && j >= 0 {
		if palindrome[i] != palindrome[j] {
			return false
		}
		i++
		j--
	}
	return true
}
