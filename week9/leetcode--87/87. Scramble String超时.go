var searched map[string][]map[string]bool

func recursion(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    if len(searched[s1]) > 0 {
        for _, Map := range searched[s1] {
            if _, ok := Map[s2]; ok {
                return Map[s2]
            }
        }
    } else {
        searched[s1] = make([]map[string]bool, 0)
    }

    Map := make(map[string]bool)
    for i:=1; i<len(s1);i++ {
        if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) || (isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i])) {
            Map[s2]=true
            searched[s1] = append(searched[s1], Map)
            return true
        }
    }
    Map[s2]=false
    searched[s1] = append(searched[s1], Map)
    return false
}

func isScramble(s1 string, s2 string) bool {
    searched = make(map[string][]map[string]bool)
    return recursion(s1, s2)
}
