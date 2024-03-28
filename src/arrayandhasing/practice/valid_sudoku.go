package practice

/*
Problem:
Ref: https://leetcode.com/problems/valid-sudoku/description/
Given:
- Sudoku board 9x9 with filled cells
- Board is partially filled
- Board can be unsolvable
- Condition:
  - Each row has no repetition
  - Each column has no repetition
  - Each the sub-box of the main grid has no repetition

Return:
- Sudoku board is valid or not

Constraints:
- board length: 9
- board[i] length: 9
- board[i][j]: [1, 9] or '.'
*/

/*
First approach:
Easiest way, we can loop every row, column and sub-box to check repetition

Time complexity: O(3n * n) ~ O(n^2) (n = 9)
Space complexity: O(n)
*/
func isValidSudoku(board [][]byte) bool {
	//Check row
	for i := 0; i < 9; i++ {
		checkArr := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			if checkArr[board[i][j]] {
				return false
			} else {
				checkArr[board[i][j]] = true
			}
		}
	}

	//Check column
	for i := 0; i < 9; i++ {
		checkArr := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			if checkArr[board[j][i]] {
				return false
			} else {
				checkArr[board[j][i]] = true
			}
		}
	}

	//Check sub-box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkArr := make(map[byte]bool)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[x+3*i][y+3*j] == '.' {
						continue
					}

					if checkArr[board[x+3*i][y+3*j]] {
						return false
					} else {
						checkArr[board[x+3*i][y+3*j]] = true
					}
				}
			}
		}
	}

	return true
}
