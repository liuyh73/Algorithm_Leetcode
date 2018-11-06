func numDistinct(s string, t string) int {
    lenS := len(s)
    lenT := len(t)
    dp := make([][]int, 0)
    for i:=0; i<=lenS; i++ {
        dp = append(dp, make([]int, lenT+1))
        dp[i][0]=1
    }

    for j:=1; j<=lenT; j++ {
        dp[0][j]=0
    }

    for i:=1; i<=lenS; i++ {
        for j:=1; j<=lenT; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }

    return dp[lenS][lenT]
}