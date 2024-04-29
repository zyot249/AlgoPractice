package practice

/*
Problem:
Ref: https://leetcode.com/problems/surrounded-regions/description/

Given:
- An m*n matrix (board) contains 'X' and 'O'
- Capture all regions that are 4-directionally surrounded by 'X'
- A region is captured by flipping all 'O's into 'X's in that surrounded region.

Return:

Constraints:
- m, n: [1, 200]
*/

/*
First approach:
Use DFS to traverse all the matrix and also trace the satisfied regions
*/

func solve(board [][]byte) {
	w, h := len(board), len(board[0])

	visited := make([][]bool, w)
	for i := 0; i < w; i++ {
		visited[i] = make([]bool, h)
	}

	var dfs func(i, j int, p *[][2]int) bool
	dfs = func(i, j int, p *[][2]int) bool {
		if i < 0 || i >= w || j < 0 || j >= h {
			return false
		}

		if visited[i][j] || board[i][j] == 'X' {
			return true
		}

		visited[i][j] = true
		cell := [2]int{i, j}
		*p = append(*p, cell)

		ok1 := dfs(i-1, j, p)
		ok2 := dfs(i+1, j, p)
		ok3 := dfs(i, j-1, p)
		ok4 := dfs(i, j+1, p)
		return ok1 && ok2 && ok3 && ok4
	}

	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			path := make([][2]int, 0)
			if dfs(x, y, &path) {
				for _, cell := range path {
					board[cell[0]][cell[1]] = 'X'
				}
			}
			clear(path)
		}
	}
}
