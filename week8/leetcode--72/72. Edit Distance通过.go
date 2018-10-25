func min(priceDelete, priceReplace, priceInsert int) int {
    if priceDelete <= priceReplace && priceDelete <= priceInsert {
        return 0
    } else if priceReplace <= priceDelete && priceReplace <= priceInsert {
        return 1
    } else if priceInsert <= priceReplace && priceInsert <= priceDelete {
        return 2
    }
    return 0
}

func minDistance(word1 string, word2 string) int {
    word1Len := len(word1)
    word2Len := len(word2)
    dp := [][]int{}
    word1 = "."+word1
    word2 = "."+word2
    for i:=0; i<=word1Len; i++ {
        slice := make([]int, word2Len+1)
        dp = append(dp, slice)
    }
    i:=1
    j:=1
    for i<word1Len && j<word2Len {
        if word1[i] == word2[j] {
            dp[i][j] = dp[i-1][j-1]
            i++
            j++
        } else {
            indexMin :=  min(dp[i][j-1], dp[i-1][j-1], dp[i-1][j])
            // delete, replace, insert
            if indexMin == 0 {
                dp[i][j] = dp[i][j-1] + 1
                i++
            } else if indexMin == 1 {
                dp[i][j] = dp[i-1][j-1] + 1
                i++
                j++
            } else {
                dp[i][j] = dp[i-1][j] + 1
                j++
            }
        }
    }
    price := 0
    if i!=word1Len {
        price = dp[i][j] + word1Len - i
    } else if j!=0 {
        price = dp[i][j] + word2Len - j
    }
    return price
}