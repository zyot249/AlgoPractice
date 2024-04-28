package practice

/*
Problem:
Ref: https://leetcode.com/problems/palindrome-partitioning/description/

Given:
- A string (s)
- partition s such that every substring of partition is a palindrome

Return:
- All possible palindrome partitioning of s

Constraints:
- length of string: [1, 16]
- string contains only lowercase English letters

Ex:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Input: s = "a"
Output: [["a"]]
*/

/*
First approach:
Firstly, I need to find the first palindrome. It starts from the start of string. (move to next char until finding a palindrome)
After finding it, I will continue to partition the rest of string.
Base condition:
If string length = 0 => empty
If string length = 1 => return it

Time complexity: O(n * 2^t) where t is the largest partitions
*/

func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		res = append(res, []string{})
		return res
	} else if len(s) == 1 {
		res = append(res, []string{s})
		return res
	}

	for i := 0; i < len(s); i++ {
		substring := s[:i+1]
		if !isPalindrome(substring) {
			continue
		}

		partitions := partition(s[i+1:])
		for _, part := range partitions {
			res = append(res, append([]string{substring}, part...))
		}
	}

	return res
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
