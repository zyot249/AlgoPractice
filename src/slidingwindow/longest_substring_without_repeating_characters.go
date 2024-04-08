package slidingwindow

/*
Problem:
Ref: https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Given:
- string s

Return:
- The length of longest substring without repeating characters

Constraints:
- string length: [0, 5*10^4]
- string contains: English letters, digits, symbols and space
*/

/*
First approach:
I will iterate all the characters of string.
For each character, I will use a hashset to check if it exists in substring and store the position I met it.
If it existed, count the number of elements in hashset and compare with max length.
Then, clear the hashset and add again all the elements after the duplicated character to the current character to set
Finally, add visiting character to set

Time complexity: O(n^2)
Space complexity: O(n)

Improvement 1:
I will store the start position of substring
If I meet a character again and its position in set >= start position of substring, I will calculate length of new substring and compare to max value.
After that, I update the start position of new substring to next position of old position I visit the duplicated character.
Then, update new position of the duplicated character.

Time complexity: O(n)
Space complexity: O(n)
*/

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	set := make(map[byte]int)
	startPos := 0
	for i := 0; i < len(s); i++ {
		foundPos, found := set[s[i]]
		if found && foundPos >= startPos {
			maxLen = max(maxLen, i-startPos)
			startPos = foundPos + 1
		}

		set[s[i]] = i
	}

	maxLen = max(maxLen, len(s)-startPos)

	return maxLen
}
