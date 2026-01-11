package neetcode150

// Problem: Design a time-based key-value store that supports:
//   - Set(key, value, timestamp)
//   - Get(key, timestamp) → returns the value with the largest timestamp ≤ given timestamp
//
// Pattern: HashMap + Binary Search (Lower Bound / Last Valid Value)
//
// Intuition:
// For each key, values are set with strictly increasing timestamps.
// This means for every key, its values form a sorted list by timestamp.
//
// When calling Get(key, timestamp):
// - We need the most recent value whose timestamp ≤ given timestamp
// - This is exactly a "rightmost valid element" problem
// - Binary Search is the optimal tool here
//
// Key Observations:
// 1️⃣ Each key maps to a list of (value, timestamp) pairs
// 2️⃣ The list is already sorted by timestamp (because Set is called in order)
// 3️⃣ We binary-search to find the latest timestamp ≤ requested timestamp
// 4️⃣ If none exists → return empty string
//
// Time Complexity:
//   Set → O(1) amortized
//   Get → O(log n) per key
//
// Space Complexity: O(n) total stored values

type TimeMap struct {
	// Maps each key to a slice of values sorted by timestamp
	KeyVal map[string][]Value
}

// Stores a value along with its timestamp
type Value struct {
	val       string
	timestamp int
}

// Constructor initializes the TimeMap
func Constructor() TimeMap {
	return TimeMap{
		KeyVal: make(map[string][]Value),
	}
}

// Set stores the value for a key at a given timestamp
func (t *TimeMap) Set(key string, value string, timestamp int) {
	// Since timestamps are strictly increasing per key,
	// we can safely append without sorting
	t.KeyVal[key] = append(
		t.KeyVal[key],
		Value{val: value, timestamp: timestamp},
	)
}

// Get returns the value with the largest timestamp ≤ given timestamp
func (t *TimeMap) Get(key string, timestamp int) string {
	arr := t.KeyVal[key]

	left, right := 0, len(arr)-1
	result := ""

	// Binary search to find the rightmost valid timestamp
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid].timestamp <= timestamp {
			// Valid candidate → store result and search right for a later one
			result = arr[mid].val
			left = mid + 1
		} else {
			// Timestamp too large → discard right half
			right = mid - 1
		}
	}

	return result
}
