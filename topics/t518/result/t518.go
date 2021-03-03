package result

/*
338. 比特位计数

给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]

示例 2:

输入: 5
输出: [0,1,1,2,1,2]

进阶:

    给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
    要求算法的空间复杂度为O(n)。
    你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。


*/

/*
思路：从0开始遍历，每次和1与运算为1则总数加1
*/
func CountBits(num int) []int {
	result, sum, j := make([]int, num+1), 0, 0

	for i := 0; i <= num; i++ {
		sum, j = 0, i
		for ; j > 0; {
			if j&1 == 1 {
				sum++
			}
			j = j >> 1
		}
		result[i] = sum
	}

	return result
}

/*
思路：last为前一个数的数量，当前数由第一位开始&1，为1则总数+1并退出，为0则总数-1，j右移一位继续
*/
func CountBits2(num int) []int {
	result, last, j := make([]int, num+1), 0, 0

	for i := 1; i <= num; i++ {
		j = i
		for {
			if j&1 == 1 {
				last++
				break
			}

			j = j >> 1
			last--
		}
		result[i] = last
	}

	return result
}
