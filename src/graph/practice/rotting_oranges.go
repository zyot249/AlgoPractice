package practice

/*
Problem:
Ref: https://leetcode.com/problems/rotting-oranges/description/

Given:
- An m*n grid
- Each cell can be:
0 => empty
1 => fresh orange
2 => rotten orange
- Every minute, any fresh orange that is 4-directionally adjacent to rotten orange becomes rotten

Return:
- Minimum number of minutes until no fresh orange exists
- If impossible => -1

Constraints:
- m, n: [1, 10]
*/

/*
First approach:
Firstly, I will find positions of all rotten oranges and count total rotten oranges and oranges
I will use BFS to find next rotten oranges from all current rotten oranges.
Stop when no more orange becomes rotten
If total rotten oranges < total oranges => impossible
*/

func orangesRotting(grid [][]int) int {
	w, h := len(grid), len(grid[0])

	totalRottenOranges, totalOranges := 0, 0
	queue := make([][2]int, 0)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if grid[x][y] == 0 {
				continue
			}

			totalOranges++

			if grid[x][y] == 2 {
				totalRottenOranges++
				queue = append(queue, [2]int{x, y})
			}
		}
	}

	var rottenOrange func(i, j int, q *[][2]int) int
	rottenOrange = func(i, j int, q *[][2]int) int {
		if i < 0 || i >= w || j < 0 || j >= h || grid[i][j] == 0 || grid[i][j] == 2 {
			return 0
		}

		grid[i][j] = 2
		*q = append(*q, [2]int{i, j})
		return 1
	}

	count := 0
	for len(queue) > 0 {
		nextQueue := make([][2]int, 0)
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			totalRottenOranges += rottenOrange(cell[0]-1, cell[1], &nextQueue)
			totalRottenOranges += rottenOrange(cell[0], cell[1]+1, &nextQueue)
			totalRottenOranges += rottenOrange(cell[0]+1, cell[1], &nextQueue)
			totalRottenOranges += rottenOrange(cell[0], cell[1]-1, &nextQueue)
		}

		if len(nextQueue) > 0 {
			count++
		}

		queue = nextQueue
	}

	if totalOranges == totalRottenOranges {
		return count
	} else {
		return -1
	}
}
