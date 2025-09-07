package neetcode150

import (
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into
// lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Pattern: Two-Pointer Technique + String Sanitization
// Approach:
// - Filter out non-alphanumeric characters.
// - Convert everything to lowercase to ignore case differences.
// - Use two pointers (start and end) to compare characters inward.
// - If all match → it's a palindrome, else return false.
// Complexity: O(n) time, O(n) space.
func isPalindrome(s string) bool {
	// Step 1: Keep only letters and digits (ignore punctuation, spaces, etc.)
	// We'll store ASCII values of valid characters in `arr`
	var arr []uint8
	for i := 0; i < len(s); i++ {
		// Check if s[i] is a digit (0–9), uppercase (A–Z), or lowercase (a–z)
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			arr = append(arr, s[i]) // keep it
		}
	}

	// Step 2: Convert arr into lowercase string
	// This ensures case-insensitivity ("A" == "a")
	palindrome := strings.ToLower(string(arr))

	// Step 3: Use Two Pointers to check palindrome
	// One pointer from the start (i), another from the end (j)
	i := 0
	j := len(palindrome) - 1

	// Compare characters while moving pointers inward
	for i < len(palindrome) && j >= 0 {
		if palindrome[i] != palindrome[j] {
			// If mismatch → not a palindrome
			return false
		}
		i++ // move forward
		j-- // move backward
	}

	// If no mismatches, it is a valid palindrome
	return true
}
