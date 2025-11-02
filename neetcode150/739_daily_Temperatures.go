package neetcode150

// Problem: Daily Temperatures
// Given an array of daily temperatures, return an array where each index `i`
// tells you how many days you’d have to wait until a warmer temperature occurs.
// If no such future day exists, put 0.
//
// Example:
// Input:  [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
//
// Pattern: Monotonic Stack (Decreasing Stack)
// -------------------------------------------
// Why? Because we are looking for the *next greater* temperature for each day.
//
// Intuition:
// - Traverse the array from **right to left** (reverse order).
// - Keep a stack of past days (temperature + index), but maintain it
//   in *strictly decreasing order* by temperature.
//   → This means the top of the stack is always the **next warmer day**.
// - For each temperature:
//     1. Pop colder or equal temperatures from the stack — they’re useless.
//     2. The element at the top (if any) gives us the next warmer temperature.
//     3. Calculate the difference in indices to get “days until warmer”.
//     4. Push the current day (temperature + index) onto the stack.
//
// Time Complexity: O(n) — each temperature is pushed/popped at most once.
// Space Complexity: O(n) — for the stack.

func dailyTemperatures(temperatures []int) []int {
	// Define a helper struct to store both temperature and its index
	type Temps struct {
		temp int
		idx  int
	}

	// Stack to store Temps in decreasing order of temperature
	var temps []Temps

	// Initialize stack with the last day's temperature
	temps = append(temps, Temps{
		temp: temperatures[len(temperatures)-1],
		idx:  len(temperatures) - 1,
	})

	// Last day has no future day → set to 0
	temperatures[len(temperatures)-1] = 0

	// Traverse from second-last day backward
	for i := len(temperatures) - 2; i >= 0; i-- {
		currentTemp := temperatures[i]
		temperatures[i] = 0 // default: assume no warmer day ahead

		// Pop all temperatures <= current (not warmer)
		for len(temps) != 0 && temps[len(temps)-1].temp <= currentTemp {
			temps = temps[:len(temps)-1]
		}

		// If stack is not empty → top of stack is next warmer day
		if len(temps) > 0 {
			temperatures[i] = temps[len(temps)-1].idx - i
		}

		// Push current day to stack for future comparisons
		temps = append(temps, Temps{
			temp: currentTemp,
			idx:  i,
		})
	}

	return temperatures
}
