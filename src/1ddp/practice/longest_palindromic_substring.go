package practice

/*
Problem:
Ref: https://leetcode.com/problems/longest-palindromic-substring/

Given:
- A string s

Return:
- The longest palindromic substring in s

Constraints:
- length of s: [1, 1000]
- s consists of only digits and English letters
*/

/*
First approach;
I will loop for all string to find all the starting palindromes in string.
They are single letters or substring of the same letter.
From that, I will find next longer palindrome by expending to both right and left side of each palindrome
Stop when finding no next palindrome

Time complexity: O(n^2)
*/

func longestPalindrome(s string) string {
	result := []int{0, 0}
	queue := make([][]int, 0)

	i := 0
	for i < len(s) {
		j := i + 1
		for j < len(s) {
			if s[i] != s[j] {
				break
			}
			j++
		}

		palindrome := []int{i, j - 1}
		if j-1-i > result[1]-result[0] {
			result = palindrome
		}

		queue = append(queue, palindrome)
		i = j
	}

	for len(queue) > 0 {
		count := len(queue)
		for _, palindrome := range queue {
			nextStart := palindrome[0] - 1
			nextEnd := palindrome[1] + 1
			if nextStart < 0 || nextEnd >= len(s) {
				continue
			}

			if s[nextStart] == s[nextEnd] {
				nextPalindrome := []int{nextStart, nextEnd}
				if nextEnd-nextStart > result[1]-result[0] {
					result = nextPalindrome
				}

				queue = append(queue, nextPalindrome)
			}
		}

		queue = queue[count:]
	}

	return s[result[0] : result[1]+1]
}
