package algorithms

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		switch {
		case target > arr[mid]:
			low = mid + 1
		case target < arr[mid]:
			high = mid - 1
		default:
			return mid
		}
	}
	return -1
}
