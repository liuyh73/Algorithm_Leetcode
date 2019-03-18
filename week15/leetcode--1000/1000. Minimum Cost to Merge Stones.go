const INT_MAX = int(^uint(0) >> 1)
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func mergeStones(stones []int, K int) int {
    sum := make([]int, 1)
    if (len(stones)-1)%(K-1) != 0 {
        return -1
    }
    dp := make([][]int, 0)
    for i:=1; i<=len(stones); i++ {
        sum = append(sum, sum[i-1] + stones[i-1])
        dp = append(dp, make([]int, len(stones)))
    }

    for length:=K; length<=len(stones); length++ {
        for i:=0; i+length<=len(stones); i++ {
            j := i+length-1
            dp[i][j]=INT_MAX
            for k:=i ;k<j; k+=K-1 {
                dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
            }
            if (j-i)%(K-1) == 0 {
                dp[i][j] += sum[j+1]-sum[i]
            }
        }
    }
    return dp[0][len(stones)-1]
}