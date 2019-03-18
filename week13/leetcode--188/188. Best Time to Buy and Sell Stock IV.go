const INT_MIN = ^int(^uint(0) >> 1)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxProfit(k int, prices []int) int {
	if k > len(prices) {
		profit := 0
		for i := 0; i < len(prices)-1; i++ {
			if prices[i+1] > prices[i] {
				profit += prices[i+1] - prices[i]
			}
		}
		return profit
	}
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	for i := 0; i <= k; i++ {
		buy[i] = INT_MIN
	}
	for _, price := range prices {
		for i := 1; i <= k; i++ {
			buy[i] = max(buy[i], sell[i-1]-price)
			sell[i] = max(sell[i], buy[i]+price)
		}
	}
	return sell[k]
}