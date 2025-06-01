package easy

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
		newArr[len(newArr)-1] = 1

		prevArr := result[i-1]
		for j, prevIter := 1, 0; j < len(newArr)-1; j++ {
			newArr[j] = prevArr[prevIter] + prevArr[prevIter+1]
			prevIter++
		}
		result[i] = newArr
	}
	return result
}
