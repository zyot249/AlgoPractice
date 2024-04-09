package slidingwindow

/*
Problem:
Ref: https://leetcode.com/problems/longest-repeating-character-replacement/description/

Given:
- a string (s)
- an integer (k)
- You can replace any character in string into other uppercase English character at most k times.

Return:
- the length of longest substring contains same letter

Constraints:
- s length: [1, 10^5]
- k: [1, len(s)]
- s contains uppercase English letter only.
*/

/*
First approach:
Consider at a start point of substring.
Condition of moving 2nd pointer to next position is:
- exist any character satisfy: endPos - startPos - character num <= k (consider only most frequency character)
If the condition is not valid -> move the start point to the next position in string.

Time complexity: O(n)
Space complexity: O(n)

Improvement 1:
Since there are only 26 uppercase English characters
--> Use array to store fre will be faster than hashmap
*/

func characterReplacement(s string, k int) int {
	//maxLen := 0
	//startPos, endPos := 0, 0
	//freMap := make(map[byte]int)
	//maxFre := 0
	//for endPos < len(s) {
	//	freMap[s[endPos]]++
	//	fre := freMap[s[endPos]]
	//	maxFre = max(maxFre, fre)
	//	if endPos-startPos-maxFre+1 <= k {
	//		maxLen = max(maxLen, endPos-startPos+1)
	//	} else {
	//		maxFre = 0
	//		freMap[s[startPos]]--
	//		for _, f := range freMap {
	//			maxFre = max(maxFre, f)
	//		}
	//
	//		startPos++
	//	}
	//
	//	endPos++
	//}
	//
	//return maxLen

	maxLen := 0
	startPos, endPos := 0, 0
	freMap := make([]int, 26)
	maxFre := 0
	for endPos < len(s) {
		freMap[s[endPos]-'A']++
		fre := freMap[s[endPos]-'A']
		maxFre = max(maxFre, fre)
		if endPos-startPos-maxFre+1 <= k {
			maxLen = max(maxLen, endPos-startPos+1)
		} else {
			maxFre = 0
			freMap[s[startPos]-'A']--
			for _, f := range freMap {
				maxFre = max(maxFre, f)
			}

			startPos++
		}

		endPos++
	}

	return maxLen
}
