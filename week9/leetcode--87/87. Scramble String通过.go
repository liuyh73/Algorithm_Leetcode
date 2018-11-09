func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    s1Slice := strings.Split(s1, "")
    s2Slice := strings.Split(s2, "")
    sort.Strings(s1Slice)
    sort.Strings(s2Slice)
    if strings.Join(s1Slice, "") != strings.Join(s2Slice, "") {
        return false
    }
    for i:=1; i<len(s1);i++ {
        if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) || (isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i])) {
            return true
        }
    }
    return false
}
