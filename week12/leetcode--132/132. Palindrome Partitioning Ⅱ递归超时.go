const INT_MAX = int(^uint(0) >> 1) 
var cutsMap map[string]int
func isPalindrome(s string) bool {
    lenS := len(s)
    for i:=0; i<lenS/2; i++ {
        if s[i]!=s[lenS-i-1] {
            return false
        }
    }
    return true
}

func getMinCuts(s string) int{
    if cutsMap[s]!=0 {
        return cutsMap[s]
    }
    if isPalindrome(s) {
        return 0
    }
    lenS := len(s)
    minCuts := INT_MAX
    for i:=1; i<lenS; i++ {
        cuts1 := getMinCuts(s[:i])
        cuts2 := getMinCuts(s[i:])
        if cuts1 + cuts2 + 1 < minCuts {
            minCuts = cuts1 + cuts2 + 1
        }
    }
    cutsMap[s] = minCuts
    return minCuts
}

func minCut(s string) int {
    cutsMap = make(map[string]int)
    return getMinCuts(s)
}