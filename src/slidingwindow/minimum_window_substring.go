package slidingwindow

/*
Problem:
Ref: https://leetcode.com/problems/minimum-window-substring/description/

Given:
- String s and t of size m and n

Return:
- Minimum length of substring in s such that every character in t (including duplicates) is included in it
- If not exist, return ""

Constraints:
- m, n: [1, 10^5]
- s, t contain uppercase and lowercase English letters only
- the answer is unique
*/

/*
First approach:
Firstly, I will count the frequency of all characters in t
I will use 2 pointers to iterate through the string s.
The left pointer will indicate the start point of substring.
The right pointer will move to next character until reaching the condition of all characters in substring have frequency >= its frequency in t.
When the right pointer stops, the left pointer will move to next character until the above condition is broke.
Before it breaks, calculate the length of substring and compare to min value, if less than min value, store the substring

Time complexity: O(m + 2n)
Space complexity: O(1)

Improvement 1:
Try to use array instead of hashmap
*/

func minWindow(s string, t string) string {
	//freS, freT := make(map[byte]int), make(map[byte]int)
	//for i := 0; i < len(t); i++ {
	//	freT[t[i]]++
	//}
	//
	//res := ""
	//matchChars := 0
	//left, right := 0, 0
	//for right < len(s) {
	//	freS[s[right]]++
	//	if freS[s[right]] == freT[s[right]] {
	//		matchChars++
	//	}
	//
	//	if matchChars == len(freT) {
	//		for left <= right {
	//			freS[s[left]]--
	//			if freS[s[left]] < freT[s[left]] {
	//				matchChars--
	//				if res == "" || len(res) > right-left+1 {
	//					res = s[left : right+1]
	//				}
	//
	//				left++
	//				break
	//			}
	//			left++
	//		}
	//	}
	//
	//	right++
	//}

	freS, freT := make([]int, 'z'-'A'+1), make([]int, 'z'-'A'+1)
	numOfDistinctLetterInT := 0
	for i := 0; i < len(t); i++ {
		index := t[i] - 'A'
		if freT[index] == 0 {
			numOfDistinctLetterInT++
		}

		freT[index]++
	}

	res := ""
	matchChars := 0
	left, right := 0, 0
	for right < len(s) {
		index := s[right] - 'A'
		freS[index]++
		if freS[index] == freT[index] {
			matchChars++
		}

		if matchChars == numOfDistinctLetterInT {
			for left <= right {
				index = s[left] - 'A'
				freS[index]--
				if freS[index] < freT[index] {
					matchChars--
					if res == "" || len(res) > right-left+1 {
						res = s[left : right+1]
					}

					left++
					break
				}
				left++
			}
		}

		right++
	}

	return res
}
