var flag map[string][]map[string]int
func convert(word1Sub string, word2Sub string) int {
    if len(flag[word1Sub]) > 0 {
        for _, maps := range flag[word1Sub] {
            if maps[word2Sub] != 0 {
                return maps[word2Sub]
            }
        }
    } else {
        flag[word1Sub] = make([]map[string]int, 0)
    }
    if word1Sub == "" && word2Sub == "" {
        return 0
    }
    if word1Sub == "" {
        return len(word2Sub)
    }
    if word2Sub == "" {
        return len(word1Sub)
    }
    if word1Sub[0] == word2Sub[0] {
        price := convert(word1Sub[1:], word2Sub[1:])
        priceMap := make(map[string]int)
        priceMap[word2Sub] = price
        flag[word1Sub] = append(flag[word1Sub], priceMap)
        return price
    } else {
        priceDelete := convert(word1Sub[1:], word2Sub)  //delete
        priceReplace := convert(word1Sub[1:], word2Sub[1:]) // replace
        priceInsert := convert(word1Sub, word2Sub[1:])  // insert
        priceMin := min(priceDelete, priceReplace, priceInsert) 
        priceMap := make(map[string]int)
        priceMap[word2Sub] = priceMin + 1
        flag[word1Sub] = append(flag[word1Sub], priceMap)
        return priceMin + 1 
    }
}

func min(priceDelete, priceReplace, priceInsert int) int {
    var priceSmaller int
    if priceDelete < priceReplace {
        priceSmaller = priceDelete
    } else {
        priceSmaller = priceReplace
    }
    if priceSmaller < priceInsert {
        return priceSmaller
    } else {
        return priceInsert
    }
}
func minDistance(word1 string, word2 string) int {
    flag = make(map[string][]map[string]int)
    return convert(word1, word2)
}