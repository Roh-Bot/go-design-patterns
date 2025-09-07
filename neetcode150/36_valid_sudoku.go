package neetcode150

// isValidSudoku checks if a given 9x9 Sudoku board is valid.
// Pattern Used: Hash Set / Boolean Table Validation
//   - Instead of using hash maps or sets, we use fixed-size boolean arrays
//     to track whether a number has already been seen in a row, column, or 3x3 box.
//   - For each cell, we check three constraints:
//     1. The number should not repeat in the same row.
//     2. The number should not repeat in the same column.
//     3. The number should not repeat in the same 3x3 sub-box.
//
// Important: This only checks validity (no duplicates). It does not solve the puzzle.
func isValidSudoku(board [][]byte) bool {
	// rows[i][num] = true means 'num' has appeared in row i
	var rows [9][10]bool
	// cols[j][num] = true means 'num' has appeared in column j
	var cols [9][10]bool
	// boxes[boxRow][boxCol][num] = true means 'num' has appeared in that 3x3 sub-box
	var boxes [3][3][10]bool // 3x3 grid of boxes, each tracking digits 1-9

	// Traverse the entire 9x9 board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Ignore empty cells (denoted by '.')
			if board[i][j] == '.' {
				continue
			}

			// Convert character '1'..'9' into integer 1..9
			num := int(board[i][j] - '0')

			// Check if the number already exists in:
			// - the current row
			// - the current column
			// - the current 3x3 box (indexed by i/3, j/3)
			if rows[i][num] || cols[j][num] || boxes[i/3][j/3][num] {
				// If it already exists in any, the Sudoku is invalid
				return false
			}

			// Otherwise, mark this number as seen
			rows[i][num] = true
			cols[j][num] = true
			boxes[i/3][j/3][num] = true
		}
	}

	// If we complete the traversal without conflicts, the Sudoku is valid
	return true
}

//CALL
//isValidSudoku([][]byte{
//		{53, 51, 46, 46, 55, 46, 46, 46, 46},
//		{54, 46, 46, 49, 57, 53, 46, 46, 46},
//		{46, 57, 56, 46, 46, 46, 46, 54, 46},
//		{56, 46, 46, 46, 54, 46, 46, 46, 51},
//		{52, 46, 46, 56, 46, 51, 46, 46, 49},
//		{55, 46, 46, 46, 50, 46, 46, 46, 54},
//		{46, 54, 46, 46, 46, 46, 50, 56, 46},
//		{46, 46, 46, 52, 49, 57, 46, 46, 53},
//		{46, 46, 46, 46, 56, 46, 46, 55, 57},
//	})
