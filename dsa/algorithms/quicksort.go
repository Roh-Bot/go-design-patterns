package algorithms

// QuickSortLomuto sorts an array using the QuickSort algorithm.
func QuickSortLomuto(arr []int, low int, high int) []int {
	if low < high {
		pi := partitionLomuto(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSortLomuto(arr, low, pi-1)
		QuickSortLomuto(arr, pi+1, high)
	}
	return arr
}

// partitionLomuto chooses the pivot using the median of three method and partitions the array.
func partitionLomuto(arr []int, low int, high int) int {
	// Find the middle index
	mid := low + (high-low)/2

	// Arrange low, mid, and high so that arr[low] <= arr[mid] <= arr[high]
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	// Swap the median value with the high index to use the existing partition logic
	arr[mid], arr[high] = arr[high], arr[mid]
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuicksortHoare(arr []int, low int, high int) {
	if low < high {
		pi := partitionHoare(arr, low, high)
		QuicksortHoare(arr, low, pi)
		QuicksortHoare(arr, pi+1, high)
	}
}

func partitionHoare(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	return j
}
