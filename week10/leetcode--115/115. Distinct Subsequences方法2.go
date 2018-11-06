func numDistinct(s string, t string) int {
    ls := len(s)
    lt := len(t)
    dp := make([]int, lt + 1)
    dp[0] = 1
    
    for i := 1; i <= ls; i++ {
        for j := lt; j >= 1; j-- {
            if s[i - 1] == t[j - 1] {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[lt]
}