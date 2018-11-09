const MAX_INT = int(^uint(0) >> 1)
func min(num1, num2 int) (minn int) {
    minn = num1
    if minn > num2 {
        minn = num2
    }
    return
}

func minimumTotal(triangle [][]int) int {
    rows := len(triangle)
    cols := 1
    for i:=1; i<rows; i++ {
        cols = len(triangle[i])
        for j:=0; j<cols; j++ {
            if j == 0 {
                triangle[i][j] += triangle[i-1][j]
            } else if j == cols - 1 {
                triangle[i][j] += triangle[i-1][j-1]
            } else {
                triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
            }
        }
    }
    minn := MAX_INT
    for j:=0; j<cols; j++ {
        if minn > triangle[rows-1][j] {
            minn = triangle[rows-1][j]
        }
    }
    return minn
}