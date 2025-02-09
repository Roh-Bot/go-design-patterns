package techniques

func SlidingWindowFixed(arr []int, k int) int {
	l := len(arr)
	if l < k {
		return -1
	}

	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	windowSum := maxSum
	for i := k; i < l; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func SlidingWindowVariable(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < n; end++ {
		if index, found := charIndexMap[s[end]]; found && index >= start {
			start = index + 1
		}
		charIndexMap[s[end]] = end
		if maxLength < end-start+1 {
			maxLength = end - start + 1
		}
	}

	return maxLength
}
