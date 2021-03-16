package result

/*
59. 螺旋矩阵 II

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

示例 2：

输入：n = 1
输出：[[1]]



提示：

    1 <= n <= 20
*/
func GenerateMatrix(n int) [][]int {
	result := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		item := make([]int, n, n)
		result = append(result, item)
	}

	num, count, i := 0, 1, 0

	for true {
		for i = num; i < n-num; i++ {
			result[num][i] = count
			count++
		}

		if count > n*n {
			break
		}

		for i = num + 1; i < n-num; i++ {
			result[i][n-num-1] = count
			count++
		}

		if count > n*n {
			break
		}

		for i = n - num - 2; i >= num; i-- {
			result[n-num-1][i] = count
			count++
		}

		if count > n*n {
			break
		}

		for i = n - num - 2; i > num; i-- {
			result[i][num] = count
			count++
		}

		if count > n*n {
			break
		}

		num++
	}

	return result
}
