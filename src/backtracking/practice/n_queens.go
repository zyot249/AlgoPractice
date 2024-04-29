package practice

/*
Problem:
Ref: https://leetcode.com/problems/n-queens/description/

Given:
- The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other
- An integer n

Return:
- All distinct solutions of N-queens puzzle
- Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

Constraints:
- n: [1, 9]
*/

/*
First approach:
In n-queens puzzle, there is exactly 1 queen in a row.
I will use brute force to find the position of next queen in the next row

*/

func solveNQueens(n int) [][]string {
	var res [][]string
	var positions []int

	var backtrack func(n, i int)
	backtrack = func(n, i int) {
		if i == n {
			representation := representPositions(n, positions)
			res = append(res, representation)
			return
		}

		for j := 0; j < n; j++ {
			if validPosition(positions, i, j) {
				positions = append(positions, j)
				backtrack(n, i+1)
				positions = positions[:len(positions)-1]
			}
		}
	}

	backtrack(n, 0)

	return res
}

func validPosition(positions []int, x int, y int) bool {
	for posX, posY := range positions {
		if y == posY || y == posY-x+posX || y == posY+x-posX {
			return false
		}
	}

	return true
}

func representPositions(n int, positions []int) []string {
	var res []string
	for _, pos := range positions {
		res = append(res, representPosition(n, pos))
	}

	return res
}

func representPosition(n, i int) string {
	var res string
	for j := 0; j < n; j++ {
		if j == i {
			res += "Q"
		} else {
			res += "."
		}
	}

	return res
}
