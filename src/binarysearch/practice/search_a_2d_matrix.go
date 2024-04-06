package practice

/*
Problem:
Ref: https://leetcode.com/problems/search-a-2d-matrix/description/

Given:
- An integer m*n matrix
- Each row is sorted in non-decreasing order
- First element in a row is greater than last element of previous row
- A target

Return:
- true if target exists and false otherwise

Constraints:
- Time complexity: O(log(m*n))
- m, n: [1, 100]
- matrix[i][j], target: [-10^4, 10^4]
*/

/*
First approach:
We need to convert the matrix into an arr by concatenate all rows of matrix.
The concatenated arr will be sorted in non-decreasing order
Use binary search to search target.

Time complexity: O(log(m*n))
Space complexity: O(1)
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2
		x, y := getIndexInMatrix(n, mid)
		if matrix[x][y] < target {
			left = mid + 1
		} else if matrix[x][y] > target {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func getIndexInMatrix(n int, index int) (int, int) {
	x := index / n
	y := index % n
	return x, y
}
