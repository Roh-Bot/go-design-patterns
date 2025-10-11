package neetcode150

//Example 1:
//
//Input: prices = [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
//Example 2:
//
//Input: prices = [7,6,4,3,1]
//Output: 0
//Explanation: In this case, no transactions are done and the max profit = 0.

// maxProfit finds the maximum profit you can achieve from a single buy and sell
// given the daily prices of a stock. You must buy before you sell.
// Approach: use two pointers (left = buy day, right = sell day),
// and slide the window to always keep track of the best profit.
func maxProfit(prices []int) int {
	left := 0  // buy day
	right := 1 // sell day
	maxProfit := 0

	for right < len(prices) {
		// if current price is less than buy price, move buy pointer
		if prices[left] > prices[right] {
			left = right
			right++
			continue
		}

		// calculate profit if sold today
		maxProfit = max(maxProfit, prices[right]-prices[left])

		// move sell pointer forward
		right++
	}

	return maxProfit
}
