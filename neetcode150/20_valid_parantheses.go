package neetcode150

// isValid checks if the input string of brackets is valid.
// A string is valid if every opening bracket has a matching closing bracket
// in the correct order.
//
// 🧠 Pattern: Stack
// We push the expected closing brackets onto the stack when we see an opener.
// When we encounter a closer, we check if it matches the top of the stack.
//
// ✅ Time Complexity: O(n) — each character is pushed/popped at most once.
// ✅ Space Complexity: O(n) — stack may hold up to n/2 elements in the worst case.
func isValid(s string) bool {
	// Edge case: odd-length strings can never be balanced
	if len(s)%2 != 0 {
		return false
	}

	// Mapping of opening → expected closing brackets
	bracks := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := new(Stack) // our custom stack

	for i := 0; i < len(s); i++ {
		char := s[i]

		// Step 1: If it's an opening bracket, push its expected closer
		if closer, ok := bracks[char]; ok {
			stack.Push(closer)
			continue
		}

		// Step 2: If it's a closing bracket,
		// check if it matches the top of the stack
		if stack.IsEmpty() || stack.Peek() != char {
			return false
		}

		// Step 3: Pop the matched closer
		stack.Pop()
	}

	// Step 4: Valid only if all openers were matched
	return stack.IsEmpty()
}

// Stack Helper Implementation
type Stack []byte

// Pop removes the top element
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	*s = (*s)[:len(*s)-1]
}

// Push adds a new element to the top
func (s *Stack) Push(data byte) {
	*s = append(*s, data)
}

// Peek returns the top element without removing it
func (s *Stack) Peek() byte {
	return (*s)[len(*s)-1]
}

// IsEmpty checks if the stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
