/****
题目描述
给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。

Consider the following matrix:
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

Given target = 5, return true.
Given target = 20, return false.
解题思路
要求时间复杂度 O(M + N)，空间复杂度 O(1)。其中 M 为行数，N 为 列数。

该二维数组中的一个数，小于它的数一定在其左边，大于它的数一定在其下边。因此，从右上角开始查找，就可以根据 target 和当前元素的大小关系来缩小查找区间，当前元素的查找区间为左下角的所有
****/
//判断是否在内部
func Find(target int, matrix [][]int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	//从右上角开始查找
	m := len(matrix)    //行
	n := len(matrix[0]) //列

	// 右上角
	org := matrix[0][n-1]
	i := 0
	j := n - 1
	for {
		if target > org {
			j--
			if j < 0 {
				return false
			}
		}
		if target < org {
			i++
			if i > m-1 {
				return false
			}
		}
		if target < org {
			return true
		}
		org = matrix[i][j]
	}
}
