package result

/*
54. 螺旋矩阵

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 10
    -100 <= matrix[i][j] <= 100

*/

/*
思路：遍历矩形最外层，剩余内圈矩形作为新的矩形递归
*/
func SpiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		result = append(result, matrix[0][i])
	}

	if m == 1 {
		return result
	}

	for i := 1; i < m; i++ {
		result = append(result, matrix[i][n-1])
	}

	if n == 1 {
		return result
	}

	for i := n - 2; i >= 0; i-- {
		result = append(result, matrix[m-1][i])
	}

	if m == 2 {
		return result
	}

	for i := m - 2; i >= 1; i-- {
		result = append(result, matrix[i][0])
	}

	if n == 2 {
		return result
	}

	leftMatrix := make([][]int, 0)

	for i := 1; i < m-1; i++ {
		item := matrix[i][1 : n-1]
		leftMatrix = append(leftMatrix, item)
	}

	result = append(result, SpiralOrder(leftMatrix)...)

	return result
}
