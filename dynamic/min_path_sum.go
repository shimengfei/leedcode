package dynamic

//TODO 空间复杂度可以进行优化
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, col := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	//i=0,j=0
	dp[0][0] = grid[0][0]
	//i>0,j=0时
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	//i=0,j>0时
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	//i>0,j>0
	for i := 1; i < rows; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
