package stackqueue

import "fmt"

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
func (s *Stack[T]) Pop() (T, bool) {
	var zeroVal T
	if s.IsEmpty() {
		return zeroVal, false
	}
	index := len(*s) - 1 // Get the index of the top element
	val := (*s)[index]   // Get the index of the top element
	*s = (*s)[:index]    // Remove it from the stack by slicing
	return val, true
}

// Peek returns the top element of the stack without removing it. Returns false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	var zeroValue T
	if s.IsEmpty() {
		return zeroValue, false
	}
	index := len(*s) - 1 // Get the index of the top element
	return (*s)[index], true
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	*s = nil
}

func main() {
	stack := new(Stack[string]) // Create a stack variable of type Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")

	stack.Push("new")
	stack.Push("elements")

	element, _ := stack.Peek()
	fmt.Printf("Peek at top element: %s", element)
	fmt.Println("Stack size:", stack.Size())

	stack.Clear()
	fmt.Println("Stack cleared. Is empty?", stack.IsEmpty())
}
