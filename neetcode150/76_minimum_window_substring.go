package neetcode150

import "math"

// Problem: Find the smallest substring in 's' that contains all characters of 't'.
// Approach:
//  1. Build a frequency map (neededMap) for all characters in 't'.
//  2. Use two pointers (left, right) to create a sliding window over 's'.
//  3. Expand 'right' until the window contains all required characters.
//  4. Once valid, move 'left' to shrink and minimize the window size.
//  5. Track the smallest valid window seen so far and return it.
// Time Complexity: O(n + m), where n = len(s), m = len(t)
// Space Complexity: O(n + m) due to maps

func minWindow(s string, t string) string {
	// ✅ Step 1: Handle edge cases
	if len(t) == 0 || len(s) == 0 || len(t) > len(s) {
		return ""
	}

	// ✅ Step 2: Create frequency map of characters needed from 't'
	neededMap := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		neededMap[t[i]]++
	}

	// ✅ Step 3: Initialize sliding window variables
	havingMap := make(map[byte]int)
	needed := len(neededMap)       // Unique chars required
	having := 0                    // Unique chars currently satisfied
	left, right := 0, 0            // Window boundaries
	start := 0                     // Start index of best window
	minWindowLength := math.MaxInt // Track smallest valid window length

	// ✅ Step 4: Expand the window using 'right' pointer
	for right < len(s) {
		currentChar := s[right]
		havingMap[currentChar]++

		// If this char completes one requirement, increase 'having'
		if neededMap[currentChar] > 0 && havingMap[currentChar] == neededMap[currentChar] {
			having++
		}

		// ✅ Step 5: Contract from the left while window is valid
		for having == needed {
			currentWindowLength := right - left + 1

			// Update result if we found a smaller valid window
			if currentWindowLength < minWindowLength {
				minWindowLength = currentWindowLength
				start = left
			}

			// Remove the leftmost character and shrink window
			leftChar := s[left]
			havingMap[leftChar]--
			left++

			// If removing this breaks validity, reduce 'having'
			if neededMap[leftChar] > 0 && havingMap[leftChar] < neededMap[leftChar] {
				having--
			}
		}

		right++ // Expand the window further
	}

	// ✅ Step 6: Return the smallest window found
	if minWindowLength == math.MaxInt {
		return ""
	}
	return s[start : start+minWindowLength]
}
