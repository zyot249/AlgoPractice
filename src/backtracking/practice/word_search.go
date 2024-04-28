package practice

/*
Problem:
Ref: https://leetcode.com/problems/word-search/description/

Given:
- A m x n grid board of characters
- A string word
- The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once

Return:
- True if word exists and false otherwise

Constraints:
- m, n: [1, 6]
- word length: [1, 15]
- board and word contain lowercase and uppercase English letters only
*/

/*
First approach:
I will find the word from each cell in the board.
If the visiting cell has the same character as the first character of word, find the rest of word from possible next cells
Store visiting cell for avoiding using 1 cell more than once

Time complexity: O(n * m * 4^t) where t is length of word

Improvement 1:
Use array for checking visited cells instead of hashmap

Improvement 2:
Instead of getting next moves with many actions with array, check condition and do immediately
*/

func exist(board [][]byte, word string) bool {
	w, h := len(board), len(board[0])
	visited := make([]bool, w*h)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(w, h, i, j, board, word, visited) {
				return true
			}
		}
	}

	return false
}

func search(width, height, x, y int, board [][]byte, word string, visitedCells []bool) bool {
	if board[x][y] != word[0] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	visitedCells[getCell(height, x, y)] = true

	if (x+1 < width && !visitedCells[getCell(height, x+1, y)] && search(width, height, x+1, y, board, word[1:], visitedCells)) ||
		(x-1 >= 0 && !visitedCells[getCell(height, x-1, y)] && search(width, height, x-1, y, board, word[1:], visitedCells)) ||
		(y+1 < height && !visitedCells[getCell(height, x, y+1)] && search(width, height, x, y+1, board, word[1:], visitedCells)) ||
		(y-1 >= 0 && !visitedCells[getCell(height, x, y-1)] && search(width, height, x, y-1, board, word[1:], visitedCells)) {
		return true
	}

	visitedCells[getCell(height, x, y)] = false

	return false
}

func getCell(height, i, j int) int {
	return i*height + j
}
