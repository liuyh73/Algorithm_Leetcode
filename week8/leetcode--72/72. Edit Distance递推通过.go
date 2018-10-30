func min(priceDelete, priceReplace, priceInsert int) int {
    if priceDelete <= priceReplace && priceDelete <= priceInsert {
        return priceDelete
    } else if priceReplace <= priceDelete && priceReplace <= priceInsert {
        return priceReplace
    } else if priceInsert <= priceReplace && priceInsert <= priceDelete {
        return priceInsert
    }
    return 0
}

func minDistance(word1 string, word2 string) int {
    word1Len := len(word1)
    word2Len := len(word2)
    word1 = "."+word1
    word2 = "."+word2
    dp := [][]int{}
    for i:=0; i<=word1Len; i++ {
        slice := make([]int, word2Len+1)
        dp = append(dp, slice)
        if i == 0 {
            for j:=0; j<=word2Len; j++ {
                dp[i][j] = j
            }
        }
        dp[i][0] =  i
    }

    for i:=1; i<word1Len+1; i++ {
        for j:=1; j<word2Len+1; j++ {
            if word1[i] == word2[j] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
            }
        }
    }

    return dp[word1Len][word2Len]
}