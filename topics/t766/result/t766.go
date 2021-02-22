package result

/*
766. 托普利茨矩阵

给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。

如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。



示例 1：

输入：matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
输出：true
解释：
在上述矩阵中, 其对角线为:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]"。
各条对角线上的所有元素均相同, 因此答案是 True 。

示例 2：

输入：matrix = [[1,2],[2,2]]
输出：false
解释：
对角线 "[1, 2]" 上的元素不同。



提示：

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 20
    0 <= matrix[i][j] <= 99

*/
func IsToeplitzMatrix(matrix [][]int) bool {
	m, n, i, j := len(matrix), len(matrix[0]), 0, 0

	for i = 0; i < n-1; i++ {
		for j = 1; j < m; j++ {
			if i+j > n-1 {
				continue
			}

			if matrix[0][i] != matrix[j][j+i] {
				return false
			}
		}
	}

	for i = 1; i < m-1; i++ {
		for j = i + 1; j < m; j++ {
			if j-i > n-1 {
				continue
			}

			if matrix[i][0] != matrix[j][j-i] {
				return false
			}
		}
	}

	return true
}
