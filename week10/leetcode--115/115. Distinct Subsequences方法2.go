func numDistinct(s string, t string) int {
    lenS := len(s)
    lenT := len(t)
    dp := make([]int, lenT+ 1)
    dp[0] = 1
    for i := 1; i <= lenS ; i++ {
        for j := lenT ; j >= 1; j-- {
            if s[i - 1] == t[j - 1] {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[lenT]
}