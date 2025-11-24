package neetcode150

import (
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	// We are applying Binary Search on the "answer" (eating speed).
	// The smallest possible speed is 1 banana/hr.
	// The largest possible speed is the maximum pile size.
	low := 1
	high := slices.Max(piles)

	// Track the minimum valid speed found so far.
	minSpeed := math.MaxInt

loop:
	// Standard binary search loop.
	for low <= high {
		// mid represents the current eating speed we are testing.
		mid := low + (high-low)/2

		// Calculate how many hours it would take to finish all piles at speed mid.
		timeTaken := 0
		for i := 0; i < len(piles); i++ {

			// Optional optimization:
			// If time already exceeds allowed h hours, mid is too slow.
			if timeTaken > h {
				// So we increase speed (shift search to the right).
				low = mid + 1
				continue loop // break out early — no need to continue calculation
			}

			// Koko needs ceil(piles[i] / mid) hours for each pile.
			// We use math.Ceil for that.
			timeTaken += max(1, int(math.Ceil(float64(piles[i])/float64(mid))))
		}

		// Now decide how to move the binary search boundaries:
		switch {
		// Case 1: Current speed mid is enough to finish within h hours.
		case timeTaken <= h:
			// This speed is valid, so try slower speeds.
			high = mid - 1
			// Track the smallest valid speed seen so far.
			minSpeed = min(minSpeed, mid)

		// Case 2: Took too long -> mid is too slow.
		case timeTaken > h:
			// Increase speed (search in the right half).
			low = mid + 1
		}
	}

	// After binary search completes, minSpeed holds the smallest valid speed.
	return minSpeed
}
