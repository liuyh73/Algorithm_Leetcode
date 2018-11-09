func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    rows := len(obstacleGrid)
    cols := len(obstacleGrid[0])
    dp := make([][]int, 0)

    for i:=0; i<=rows; i++ {
        row := make([]int, cols+1)
        dp = append(dp, row)
    }
    dp[1][1] = 1
    if obstacleGrid[0][0] == 1 {
        dp[1][1] = 0
    }
    for i:=1; i<=rows; i++ {
        for j:=1; j<=cols; j++ {
            if obstacleGrid[i-1][j-1] == 0 && !(i==1 && j==1) {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[rows][cols]
}