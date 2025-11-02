package neetcode150

import "strconv"

// Problem: Evaluate a mathematical expression written in Reverse Polish Notation (RPN).
// Example: ["2","1","+","3","*"] → ((2 + 1) * 3) = 9
//
// Approach (Stack-based Evaluation):
//  1. Iterate through each token in the list.
//  2. If the token is a number, push it onto the stack.
//  3. If the token is an operator (+, -, *, /):
//      - Pop the top two numbers from the stack.
//      - Apply the operation (order matters: first popped is second operand).
//      - Push the result back onto the stack.
//  4. After processing all tokens, the stack will contain exactly one element — the final answer.
//
// Time Complexity: O(n), where n = number of tokens
// Space Complexity: O(n) in the worst case (if all tokens are numbers)

func evalRPN(tokens []string) int {
	stack := make(Stack[int], 0)

	// ✅ Step 1: Traverse through all tokens
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// ✅ Step 2: If token is a number, push to stack
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
			continue
		}

		// ✅ Step 3: Otherwise, token is an operator
		// Pop top two numbers (note: second popped is actually the right operand)
		secondNum := stack.Pop()
		firstNum := stack.Pop()

		result := 0
		switch token {
		case "+":
			result = firstNum + secondNum
		case "-":
			result = firstNum - secondNum
		case "*":
			result = firstNum * secondNum
		case "/":
			result = firstNum / secondNum
		}

		// ✅ Step 4: Push the result back for further evaluation
		stack.Push(result)
	}

	// ✅ Step 5: Final value left in the stack is the answer
	return stack.Pop()
}
