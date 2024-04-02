package practice

import "strings"

/*
Problem:
Ref: https://leetcode.com/problems/valid-palindrome/description/

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given:
- a string

Return:
- true if the string is palindrome, and false otherwise

Constraints:
- string length: [0, 2*10^5]
- string consists only of printable ASCII characters.
*/

/*
First approach:
First of all, we will convert all the uppercase letters to lowercase and remove all non-alphanumeric characters
Then, we iterate from the start and the end at the same time and compare the character until 2 pointers meet each other
*/

func isAlphanumericCharacter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z')
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for j > i {
		for i < j && !isAlphanumericCharacter(s[i]) {
			i++
		}

		for j > i && !isAlphanumericCharacter(s[j]) {
			j--
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
