func isInterleave(s1 string, s2 string, s3 string) bool {
    s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
    if s1Len + s2Len != s3Len {
        return false
    }
    dp := make([][]bool, 0)
    for i:=0; i<=s1Len; i++ {
        slice := make([]bool, s2Len+1)
        dp = append(dp, slice)
    }
    // 初始化
    for i:=1; i<=s1Len; i++ {
        if s1[i-1] == s3[i-1] {
            dp[i][0] = true
        } else {
            break
        }
    }
    for i:=1; i<=s2Len; i++ {
        if s2[i-1] == s3[i-1] {
            dp[0][i] = true
        } else {
            break
        }
    }

    for i:=1; i<=s1Len; i++ {
        for j:=1; j<=s2Len; j++ {
            if (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
                dp[i][j] = true
            }
        }
    }

    return dp[s1Len][s2Len]
}