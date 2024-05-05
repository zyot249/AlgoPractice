package practice

import "strconv"

/*
Problem:
Ref: https://leetcode.com/problems/decode-ways/description/

Given:
- A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
- A string s containing only digits

Return:
- The number of ways to decode it

Constraints:
- length of s: [1, 100]
- s contains only digits and may contain leading zero(s)
*/

/*
First approach:
I will calculate the number of ways recursively by decoding each letter.
I also cache the result for reusing.

Time complexity: O(n)
Space complexity: O(n)
*/

func numDecodings(s string) int {
	cache := make([]int, len(s)+1)

	var decode func(string) int
	decode = func(str string) int {
		if str == "" {
			return 1
		}

		if str[0] == '0' {
			return 0
		}

		result := 0
		if cache[len(str)-1] != 0 {
			result += cache[len(str)-1]
		} else {
			result1 := decode(str[1:])
			cache[len(str)-1] = result1
			result += result1
		}

		if len(str) > 1 {
			num, _ := strconv.Atoi(str[0:2])
			if num <= 26 {
				if cache[len(str)-2] != 0 {
					result += cache[len(str)-2]
				} else {
					result1 := decode(str[2:])
					cache[len(str)-2] = result1
					result += result1
				}
			}
		}

		return result
	}

	return decode(s)
}
