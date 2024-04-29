package practice

/*
Problem:
Ref: https://leetcode.com/problems/max-area-of-island/description/

Given:
- A m * n binary matrix grid represents a map of land (1) and water (0)
- An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
- The area of an island is the number of cells with a value '1' in the island.

Return:
- The maximum area of an island
- If no island => 0

Constraints:
- m, n: [1, 50]
*/

/*
First approach:
I will form an island from a land in map.
After forming an island, store max area
I also mark visited lands for avoid duplication
*/

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	w, h := len(grid), len(grid[0])
	visited := make([][]bool, w)
	for i := 0; i < w; i++ {
		visited[i] = make([]bool, h)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= w || j < 0 || j >= h || visited[i][j] || grid[i][j] == 0 {
			return 0
		}

		visited[i][j] = true
		return 1 + dfs(i-1, j) + dfs(i, j+1) + dfs(i+1, j) + dfs(i, j-1)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				area := dfs(i, j)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
