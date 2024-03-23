package practice

import (
	"sort"
	"strings"
)

/*
Problem:
Ref: https://leetcode.com/problems/valid-anagram/
Require: Check if t is an anagram of s
Concept:
- An anagram is word or phrase formed by rearranging all the letters of a different word or phrase
Constraints:
- 1 <= string length <= 5 * 10^4
- string contains only lowercase letters
*/

/*
First approach:
if s and t differ in length => false
We can sort 2 strings.
After that, we loop over all the characters of string and compare with the rest.
Time complexity: O(nlogn + nlogn + n) -> O(nlogn)
Space complexity: O(n) (O(1) when optimizing to sort s and t directly if programming language supports sorting a string)
*/
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var ss = strings.Split(s, "")
	var ts = strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(ts)

	for i := 0; i < len(s); i++ {
		if ss[i] != ts[i] {
			return false
		}
	}

	return true
}

/*
Second approach:
If s and t differ in length => false
We can check the frequency of each character in each string
And then compare their frequency
Time complexity: O(n + n + n) -> O(n)
Space complexity: O(n + n) -> O(n)

Improvement 1:
Use only 1 map for counting the frequency difference of characters in string s and t
If there are any character having frequency != 0 -> false
Time complexity: O(n + n + n) -> O(n)
Space complexity: O(n)

Improvement 2:
Instead of using hashmap, we can use an array of size 26 (since the strings consist of lowercase English letters)
Access an index in array is faster than hashmap
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//var fs = make(map[byte]int)
	//var ft = make(map[byte]int)
	//
	//for i := 0; i < len(s); i++ {
	//	f, exist := fs[s[i]]
	//	if exist {
	//		fs[s[i]] = f + 1
	//	} else {
	//		fs[s[i]] = 1
	//	}
	//
	//	f, exist = ft[t[i]]
	//	if exist {
	//		ft[t[i]] = f + 1
	//	} else {
	//		ft[t[i]] = 1
	//	}
	//}
	//
	//for c, f := range fs {
	//	if f != ft[c] {
	//		return false
	//	}
	//}

	//var f = make(map[byte]int)
	//
	//for i := 0; i < len(s); i++ {
	//	f[s[i]]++
	//	f[t[i]]--
	//}
	//
	//for _, num := range f {
	//	if num != 0 {
	//		return false
	//	}
	//}

	var f = make([]int, 26)

	for i := 0; i < len(s); i++ {
		f[s[i]-'a']++
		f[t[i]-'a']--
	}

	for _, num := range f {
		if num != 0 {
			return false
		}
	}

	return true
}
