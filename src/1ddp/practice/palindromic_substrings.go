package practice

/*
Problem:
Ref: https://leetcode.com/problems/palindromic-substrings/description/

Given:
- A string s

Return:
- The number of palindromic substrings

Constraints:
- string s consists of lowercase English letters
- length of string: [1, 1000]
*/

/*
First approach;
To find next palindrome from current palindrome, I will extend it to both left and right side.
I will handle for case of odd and even length

Time complexity: O(n^2)
*/

func countSubstrings(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		//odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}

		//even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			result++
			l--
			r++
		}
	}

	return result
}
