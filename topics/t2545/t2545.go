package t2545

func SortTheStudents(score [][]int, k int) [][]int {
	return score
}

// 冒泡
func SortTheStudents1(score [][]int, k int) [][]int {
	m := len(score)

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if score[i][k] < score[j][k] {
				score[i], score[j] = score[j], score[i]
			}
		}
	}

	return score
}

// 折半
func SortTheStudents2(score [][]int, k int) [][]int {
	m := len(score)

	for i := 1; i < m; i++ {
		left := 0
		right := i - 1
		temp := score[i]

		for left <= right {
			mid := (left + right) / 2

			if score[i][k] < score[mid][k] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		for j := i; j >= left+1; j-- {
			score[j] = score[j-1]
		}

		score[left] = temp
	}

	return score
}

// 快速排序
func partition(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}
