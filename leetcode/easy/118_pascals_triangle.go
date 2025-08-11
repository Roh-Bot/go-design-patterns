package easy

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		newArr := make([]int, i+1)
		newArr[0] = 1
		newArr[i] = 1

		prevArr := result[i-1]
		for j, prevArrIter := 1, 0; j < i; j++ {
			newArr[j] = prevArr[prevArrIter] + prevArr[prevArrIter+1]
			prevArrIter++
		}
		result[i] = newArr
	}
	return result
}
