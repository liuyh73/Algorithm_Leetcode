const INT_MAX = int(^uint(0) >> 1) 
func minCut(s string) int {
    lenS := len(s)
    dp := make([][]int, lenS)
    minCuts := make([]int, lenS+10)
    for i:=0; i<lenS; i++ {
        dp[i] = make([]int, lenS)
        dp[i][i] = 1
        minCuts[i+1] = INT_MAX
    }
    minCuts[0] = -1
    minCuts[1] = 0
    for j:=0; j<lenS; j++ {
        for i:=0; i<lenS; i++ {
            if j > i {
                if dp[i+1][j-1]==1 && s[i] == s[j] {
                    dp[i][j] = 1
                } else if i+1==j && s[i] == s[j] {
                    dp[i][j] = 1
                }
            }
        }
    }

    for i:=1; i<=lenS; i++ {
        for j:=1; j<=i; j++ {
            if dp[j-1][i-1] == 1 && minCuts[i] > minCuts[j-1] + 1 {
                minCuts[i] = minCuts[j-1] + 1
            }
        }
    }

    return minCuts[lenS]
}