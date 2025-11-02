package neetcode150

// Stack represents a stack of strings.
type Stack[T comparable] []T

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds a new value onto the stack.
func (s *Stack[T]) Push(data T) {
	*s = append(*s, data) // Append the new value to the end of the stack
}

// Pop removes and returns the top element of the stack. Returns false if the stack is empty.
func (s *Stack[T]) Pop() T {
	index := len(*s) - 1 // Get the index of the top element
	val := (*s)[index]   // Get the index of the top element
	*s = (*s)[:index]    // Remove it from the stack by slicing
	return val
}

// Peek returns the top element of the stack without removing it. Returns false if the stack is empty.
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1] // Get the index of the top element
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	*s = nil
}
