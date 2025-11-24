package neetcode150

// Problem: Largest Rectangle in Histogram
// ---------------------------------------
// Given an array of bar heights, find the largest rectangular area that can be
// formed inside the histogram.
//
// Example:
// heights = [2,1,5,6,2,3]
// Largest rectangle = 10  (bars 5 and 6 → width=2, height=5)
//
// Pattern Used: Monotonic Stack (Increasing Stack)
// ------------------------------------------------
// Why this pattern?
// We want to know, for each bar, how FAR LEFT and FAR RIGHT it can extend
// before encountering a bar shorter than itself.
// Those two boundaries define its maximum possible rectangle.
//
// Intuition:
// - We use a stack of bars (value + starting index).
// - The stack is maintained in **increasing order of heights**.
// - When we see a bar shorter than the top of the stack:
//     → That means the rectangle of the taller bar must END here.
//     → We pop it, compute area using:
//         height = popped bar
//         width  = (current index - popped.startIndex)
// - When pushing a bar, we may adjust its "start index" backwards so that
//   it represents the leftmost boundary where this height could extend.
//
// After the loop, any bars left in the stack extend to the END of the
// histogram, so we compute areas with right boundary = last index.
//
// Time Complexity: O(n) because each bar is pushed/popped at most once.
// Space Complexity: O(n) for the stack.

type rect struct {
	index int
	val   int
}

func largestRectangleArea(heights []int) int {
	largestArea := 0
	stack := make(Stack[rect], 0)

	// Iterate over each bar in the histogram
	for i := 0; i < len(heights); i++ {
		startIdx := i

		// While stack contains bars taller than current bar,
		// they cannot extend past index i → compute their areas.
		for !stack.IsEmpty() && stack.Peek().val > heights[i] {
			elem := stack.Pop()

			// `elem.index` is the LEFTMOST index where this height could extend
			// `i` is the first index on the right where it is blocked
			startIdx = elem.index
			largestArea = max(largestArea, (i-startIdx)*elem.val)
		}

		// Push the new bar but with the earliest left boundary it can extend to.
		stack.Push(rect{index: startIdx, val: heights[i]})
	}

	// After processing all bars, remaining ones can extend all the way to the end.
	furthest := len(heights) - 1

	for !stack.IsEmpty() {
		elem := stack.Pop()
		// Width here is from elem.index to last index
		largestArea = max(largestArea, (furthest-elem.index+1)*elem.val)
	}

	return largestArea
}
