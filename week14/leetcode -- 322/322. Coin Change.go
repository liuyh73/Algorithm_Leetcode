const INT_MAX = int(^uint(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = INT_MAX
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && uint(dp[i-coins[j]]) < uint(INT_MAX) {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if uint(dp[amount]) >= uint(INT_MAX) {
		return -1
	}
	return dp[amount]
}