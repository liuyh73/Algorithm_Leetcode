func convert(word1Sub string, word2Sub string) int {
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
        return convert(word1Sub[1:], word2Sub[1:])
    } else {
        priceDelete := convert(word1Sub[1:], word2Sub)	//delete
        priceReplace := convert(word1Sub[1:], word2Sub[1:])	// replace
        priceInsert := convert(word1Sub, word2Sub[1:])	// insert
        priceMin := min(priceDelete, priceReplace, priceInsert) 
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
    return convert(word1, word2)
}