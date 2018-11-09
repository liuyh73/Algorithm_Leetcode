const MAX_INT = int(^uint(0) >> 1)
func min(num1, num2 int) (minn int) {
    minn = num1
    if minn > num2 {
        minn = num2
    }
    return
}

func minPathSum(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    dp := make([][]int, 0)

    for i:=0; i<=rows; i++ {
        row := make([]int, cols+1)
        dp = append(dp, row)
        for j:=0; j<=cols; j++ {
            dp[i][j] = MAX_INT
        }
    }
    dp[0][1] = 0
    dp[1][0] = 0
    for i:=1; i<=rows; i++ {
        for j:=1; j<=cols; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
        }
    }

    return dp[rows][cols]
}