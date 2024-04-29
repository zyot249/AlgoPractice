package practice

/*
Problem:
Ref: https://leetcode.com/problems/number-of-islands/description/

Given:
- An m * n 2d binary grid which represents a map of land ('1') and water ('0')
- An island is surrounded by water and formed by connecting adjacent lands horizontally and vertically
- 4 edges of grid are surrounded by water

Return:
- The number of islands

Constraints:
- m, n: [1, 300]

Ex:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/

/*
First approach:
I will form an island from a land in grid.
From the starting land, I use DFS/BFS to find all adjacent lands.
I also mark visited lands for finding new starting land.

Time complexity: (4^(m*n))
*/

func numIslands(grid [][]byte) int {
	w, h := len(grid), len(grid[0])
	count := 0
	visited := make([][]bool, w)
	for i := 0; i < w; i++ {
		visited[i] = make([]bool, h)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= w || j < 0 || j >= h || grid[i][j] == '0' || visited[i][j] {
			return
		}

		visited[i][j] = true
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if grid[x][y] == '1' && !visited[x][y] {
				count++
				dfs(x, y)
			}
		}
	}

	return count
}
