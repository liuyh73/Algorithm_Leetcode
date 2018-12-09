const INT_MIN = ^int(^uint(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func calculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	// 右下角
	if dungeon[rows-1][cols-1] > 0 {
		dp[rows-1][cols-1] = 0
	} else {
		dp[rows-1][cols-1] = -dungeon[rows-1][cols-1]
	}

	// 边界-右
	for i := rows - 2; i >= 0; i-- {
		if dungeon[i][cols-1] > dp[i+1][cols-1] {
			dp[i][cols-1] = 0
		} else {
			dp[i][cols-1] = dp[i+1][cols-1] - dungeon[i][cols-1]
		}
	}

	// 边界-下
	for j := cols - 2; j >= 0; j-- {
		if dungeon[rows-1][j] > dp[rows-1][j+1] {
			dp[rows-1][j] = 0
		} else {
			dp[rows-1][j] = dp[rows-1][j+1] - dungeon[rows-1][j]
		}
	}

	for i := rows - 2; i >= 0; i-- {
		for j := cols - 2; j >= 0; j-- {
			// Right
			if dungeon[i][j] > dp[i+1][j] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i+1][j] - dungeon[i][j]
			}

			// Down
			if dungeon[i][j] > dp[i][j+1] {
				dp[i][j] = min(0, dp[i][j])
			} else {
				dp[i][j] = min(dp[i][j+1]-dungeon[i][j], dp[i][j])
			}
		}
	}

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {
	// 		fmt.Print(dp[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	return dp[0][0] + 1
}