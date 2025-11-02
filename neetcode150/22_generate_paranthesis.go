package neetcode150

// Problem: Generate all combinations of well-formed parentheses for a given n pairs.
// Example: n = 3 → ["((()))","(()())","(())()","()(())","()()()"]
//
// Pattern: Backtracking (DFS with pruning)
// Idea: At any step, we can add either '(' or ')' to our growing string,
//       but we must ensure we never form an invalid sequence.
//       The recursion explores all valid configurations.
//
// Approach:
//  1. Use recursion (backtracking) with two counters: `opening` and `closing`.
//     - `opening` = how many '(' we can still add.
//     - `closing` = how many ')' we can still add.
//  2. Constraints for validity:
//     - Never let `closing < opening` → means too many ')' before '('.
//     - Never let `opening < 0` or `closing < 0` → out of bounds.
//  3. When both counters hit 0 → we’ve formed a valid combination.
//  4. Append it to result.
//  5. Explore both possibilities:
//     - Add '(' (if any left)
//     - Add ')' (if doing so won’t break validity)
//
// Time Complexity: O(2^(2n)) theoretical, but only valid combinations are built → ~O(4^n / √n) (Catalan number growth)
// Space Complexity: O(n) recursion depth (each combination of length 2n)

func generateParenthesis(n int) []string {
	var ret []string

	var backtrack func(opening, closing int, combination []byte)
	backtrack = func(opening, closing int, combination []byte) {
		// ❌ Invalid cases: no brackets left to use or more ')' than '(' so far
		if opening < 0 || closing < 0 || closing < opening {
			return
		}

		// ✅ Base case: all brackets used up → valid combination found
		if opening == 0 && closing == 0 {
			ret = append(ret, string(combination))
			return
		}

		// ✅ Choice 1: Add '(' if we still have some left
		if opening > 0 {
			backtrack(opening-1, closing, append(combination, '('))
		}

		// ✅ Choice 2: Add ')' if it won’t break validity
		if closing > opening {
			backtrack(opening, closing-1, append(combination, ')'))
		}
	}

	// ✅ Start recursion with n opening and n closing brackets available
	backtrack(n, n, nil)
	return ret
}
