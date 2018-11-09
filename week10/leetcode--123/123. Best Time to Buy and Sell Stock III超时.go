func max(a, b, c int) (maxx int) {
    maxx = a
    if maxx < b {
        maxx = b
    }
    if maxx < c {
        maxx =c
    }
    return
}
func maxProfit(prices []int) int {
    days := len(prices)
    dp := make([][]int, 0)
    for i:=0; i<=days; i++ {
        dp = append(dp, make([]int, 3))
    }

    for i:=1; i<=days; i++ {
        dp[i][1] = dp[i-1][1]
        for j:=1; j<i; j++ {
            if prices[i-1] - prices[j-1] > dp[i][1] {
                dp[i][1] = prices[i-1] - prices[j-1]
            }

            for k:=j+1; k<i; k++ {
                if dp[j][1] + prices[i-1] - prices[k-1] > dp[i][2] {
                    dp[i][2] = dp[j][1] + prices[i-1] - prices[k-1]
                }
            }
            dp[i][2] = max(dp[i][2], dp[i-1][2], dp[i][1])
        }
        // fmt.Println(dp[i][1], dp[i][2])
    }
    return dp[days][2]
}
