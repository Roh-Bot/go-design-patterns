package easy

type Stack[T comparable] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(str T) {
	*s = append(*s, str)
}

func (s *Stack[T]) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	index := len(*s) - 1
	*s = (*s)[:index]
	return true
}

func isValid(s string) bool {
	if len(s)%2 != 0 || len(s) < 2 {
		return false
	}
	bracks := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := Stack[rune]{}

	for _, v := range s {
		if _, ok := bracks[v]; ok {
			stack = append(stack, v)
		} else if len(stack) != 0 && bracks[stack[len(stack)-1]] != v {
			stack.Pop()
		} else {
			return false
		}
	}
	return len(stack) == 0
}
