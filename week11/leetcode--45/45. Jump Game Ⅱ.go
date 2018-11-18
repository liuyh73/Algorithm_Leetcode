const INT_MAX = int(^uint(0) >> 1)
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func jump(nums []int) int {
    length := len(nums)
    dp := make([]int, length)
    for i:=1; i<length; i++ {
        if i <= nums[0] {
            dp[i] = 1
        } else {
            dp[i] = INT_MAX
        }
    }

    for i:=1; i<length; i++ {
        dp[i] = min(dp[i], dp[i-1] + 1)
        for j:=1; j<= nums[i]; j++ {
            if i+j >= length {
                break
            }
            dp[i+j] = min(dp[i+j], dp[i] + 1)
        }
    }

    return dp[length-1]
}