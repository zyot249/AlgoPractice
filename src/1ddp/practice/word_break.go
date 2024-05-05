package practice

/*
Problem:
Ref: https://leetcode.com/problems/word-break/description/

Given:
- A string (s)
- A dictionary of strings (wordDict)

Return:
- True if s can be segmented into a space-separated sequence of one or more dictionary words
- False otherwise

Constraints:
- length of s: [1, 300]
- length of wordDict: [1, 1000]
- length of words in wordDict: [1, 20]
- s and wordDict[i] consist of only lowercase English letters
- All words in wordDict are unique
*/

/*
First approach:
Firstly, I find all the possible first dictionary word at the start of string.
Do recursively with the rest of string.
Cache result for reusing

Time complexity: O(n * m * k)
where: n: length of s, m: length of dictionary, k: length of word in dictionary
Space complexity: O(n)
*/

func wordBreak(s string, wordDict []string) bool {
	cache := make([]int, len(s)+1)

	var separate func(string) bool
	separate = func(str string) bool {
		if str == "" {
			return true
		}

		if cache[len(str)] == 1 {
			return false
		} else if cache[len(str)] == 2 {
			return true
		} else {
			for _, word := range wordDict {
				if len(str) < len(word) {
					continue
				}

				canSeparate := true
				for i := 0; i < len(word); i++ {
					if str[i] != word[i] {
						canSeparate = false
						break
					}
				}

				if !canSeparate {
					continue
				}

				result := separate(str[len(word):])
				if result {
					cache[len(str)] = 2
					return true
				} else {
					cache[len(str)] = 1
				}
			}
		}

		return false
	}

	return separate(s)
}
