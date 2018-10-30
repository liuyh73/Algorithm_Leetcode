func longestValidParentheses(s string) int {
    stack := make([]byte, 0)
    dp := make([]int, len(s)+1)
    max := 0
    if s=="" {
        return 0
    }
    stack = append(stack, s[0])
    for i:=1; i<len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, s[i])
        } else {
            stackLen := len(stack)
            if stackLen > 0 && stack[stackLen-1] == '(' {
                stack = stack[:stackLen-1]
                dp[stackLen-1] = dp[stackLen-1] + dp[stackLen] + 2
                dp[stackLen] = 0
            } else {
                stack = append(stack, s[i])
            }
        }
    }

    for i:=0; i<len(s)+1; i++ {
        if max < dp[i] {
            max = dp[i]
        }
    }
    return max
}