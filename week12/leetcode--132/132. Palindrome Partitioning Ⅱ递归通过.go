const INT_MAX = int(^uint(0) >> 1) 
var cutsMap map[string]int
var dp [][]int

func getMinCuts(i, j int, s string) int{
    if cutsMap[s[i:j+1]]!=0 {
        return cutsMap[s[i:j+1]]
    }
    if dp[i][j] == 1 {
        return 0
    }
    minCuts := INT_MAX
    for k:=i+1; k<=j; k++ {
        if dp[i][k-1] == 1 {
            cuts := getMinCuts(k, j, s)
            if cuts + 1 < minCuts {
                minCuts = cuts + 1
            }
        }
    }
    cutsMap[s[i:j+1]] = minCuts
    return minCuts
}

func minCut(s string) int {
    cutsMap = make(map[string]int)
    lenS := len(s)
    dp = make([][]int, lenS)
    for i:=0; i<lenS; i++ {
        dp[i] = make([]int, lenS)
        dp[i][i] = 1
    }
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
    return getMinCuts(0, lenS-1, s)
}