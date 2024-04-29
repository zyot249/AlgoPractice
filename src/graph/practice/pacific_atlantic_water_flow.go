package practice

/*
Problem:
Ref: https://leetcode.com/problems/pacific-atlantic-water-flow/description/

Given:
- An m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean
- The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges
- The island is partitioned into a grid of square cells
- An m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c)
- The rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height
Water can flow from any cell adjacent to an ocean into the ocean

Return:
- Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans

Constraints:
- m, n: [1, 200]
- heights[i][j]: [0, 10^5]
*/

/*
Solution:
From atlantic and pacific, we find all the cells rain can flow down

Time complexity: O(m*n)
*/

func pacificAtlantic(heights [][]int) [][]int {
	w, h := len(heights), len(heights[0])

	var dfs func(int, int, int, map[[2]int]bool)
	dfs = func(i, j, lastHeight int, m map[[2]int]bool) {
		if i < 0 || j < 0 || i >= w || j >= h || m[[2]int{i, j}] || heights[i][j] < lastHeight {
			return
		}

		cell := [2]int{i, j}
		m[cell] = true
		dfs(i-1, j, heights[i][j], m)
		dfs(i, j+1, heights[i][j], m)
		dfs(i+1, j, heights[i][j], m)
		dfs(i, j-1, heights[i][j], m)
	}

	toPacific := make(map[[2]int]bool)
	for x := 0; x < w; x++ {
		dfs(x, 0, heights[x][0], toPacific)
	}
	for y := 0; y < h; y++ {
		dfs(0, y, heights[0][y], toPacific)
	}

	toAtlantic := make(map[[2]int]bool)
	for x := 0; x < w; x++ {
		dfs(x, h-1, heights[x][h-1], toAtlantic)
	}
	for y := 0; y < h; y++ {
		dfs(w-1, y, heights[w-1][y], toAtlantic)
	}

	var result [][]int
	for cell, _ := range toPacific {
		if toAtlantic[cell] {
			result = append(result, []int{cell[0], cell[1]})
		}
	}

	return result
}
