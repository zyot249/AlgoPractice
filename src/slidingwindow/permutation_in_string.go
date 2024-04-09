package slidingwindow

/*
Problem:
Ref: https://leetcode.com/problems/permutation-in-string/description/

Given:
- string s1 and s2

Return:
- True if s2 contains a permutation of s1
- False otherwise

Constraints:
- 1 <= s1.length, s2.length <= 104
- s1 and s2 consist of lowercase English letters.
*/

/*
First approach:
Firstly, I count the frequency of all characters in string s1.
I iterate all the characters in string s2.
If visiting character is not in s1 -> move 1st pointer to next character, reset frequency counter.
If visiting character is in s1, I count the frequency of it:
- If its frequency in s2 < its frequency in s1 -> move 2nd pointer to next character
- If its frequency in s2 = its frequency in s1 -> if size of window < length of s1 => move 2nd pointer to next character, if equal => true
- If its frequency in s2 > its frequency in s1 -> move 1st pointer to next character and decrease its frequency until decreasing the character at 2nd pointer

Time complexity: O(m + n) m: len(s1), n: len(s2)
Space complexity: O(2m)
*/

func checkInclusion1(s1 string, s2 string) bool {
	s1Fre := make([]int, 26)
	s2Fre := make([]int, 26)
	for _, c := range s1 {
		s1Fre[c-'a']++
	}

	start := 0
	for end := 0; end < len(s2); end++ {
		freInS1 := s1Fre[s2[end]-'a']
		if freInS1 <= 0 {
			clear(s2Fre)
			start = end + 1
		} else {
			s2Fre[s2[end]-'a']++
			freInS2 := s2Fre[s2[end]-'a']
			if freInS2 == freInS1 && end-start+1 == len(s1) {
				return true
			} else if freInS2 > freInS1 {
				for start <= end {
					s2Fre[s2[start]-'a']--
					start++
					if s2[start-1] == s2[end] {
						break
					}
				}
			}
		}
	}

	return false
}

/*
Solution:
Since the size of s1 is fixed --> Fixed size sliding window
When moving window to next position, we will remove left most char and add right most char to frequency counter and re-calculate the number of matched chars
If number of matched chars is 26 => return true
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Fre := make([]int, 26)
	s2Fre := make([]int, 26)
	//initialize window
	for i := 0; i < len(s1); i++ {
		s1Fre[s1[i]-'a']++
		s2Fre[s2[i]-'a']++
	}

	matchChars := 0
	for i := 0; i < len(s1Fre); i++ {
		if s1Fre[i] == s2Fre[i] {
			matchChars++
		}
	}

	if matchChars == 26 {
		return true
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		index := s2[i] - 'a'
		s2Fre[index]--
		if s2Fre[index] == s1Fre[index] {
			matchChars++
		} else if s2Fre[index] == s1Fre[index]-1 {
			matchChars--
		}

		index = s2[i+len(s1)] - 'a'
		s2Fre[index]++
		if s2Fre[index] == s1Fre[index] {
			matchChars++
		} else if s2Fre[index] == s1Fre[index]+1 {
			matchChars--
		}

		if matchChars == 26 {
			return true
		}
	}

	return false
}
