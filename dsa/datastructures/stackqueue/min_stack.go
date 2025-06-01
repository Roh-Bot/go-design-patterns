package stackqueue

type element struct {
	val int
	min int
}

type MinStack []element

func (s *MinStack) Push(x int) {
	minimum := x
	if len(*s) > 0 && (*s)[len(*s)-1].min < x {
		minimum = (*s)[len(*s)-1].min
	}
	*s = append(*s, element{val: x, min: minimum})
}

func (s *MinStack) Pop() {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	*s = (*s)[:len(*s)-1]
}

func (s *MinStack) Top() int {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	return (*s)[len(*s)-1].val
}

func (s *MinStack) GetMin() int {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	return (*s)[len(*s)-1].min
}
