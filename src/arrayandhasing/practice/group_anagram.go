package practice

/*
Problem:
Ref: https://leetcode.com/problems/group-anagrams/
Given arr of strings
=> Group the anagrams together

Constraints:
- string length: [0, 100]
- arr length: [0, 10^4]
- string contains lowercase English letters only
*/

/*
First approach:
Find a way to convert a string into the form that can be unique for an anagram group
- 1, sort string
We visit all the elements of array, convert them into above form and group them using hashmap
Time complexity: O(n*mlogm + n) ~ O(n) (since m is small)
Space complexity: O(n)

Improvement 1:
Try to find another way of converting a string into unique key for an anagram
Instead of sort string, we convert string into arr of char frequency and use it as key
Time complexity: O(n*m + n) ~ O(n) (since m is small)
Space complexity: O(n)
*/
func groupAnagrams(strs []string) [][]string {
	//group := make(map[string][]string)
	//for _, str := range strs {
	//	runes := []rune(str)
	//	slices.Sort(runes)
	//	key := string(runes)
	//	group[key] = append(group[key], str)
	//}

	var group = make(map[[26]uint8][]string)
	for _, str := range strs {
		var key [26]uint8
		for _, c := range str {
			key[c-'a']++
		}
		group[key] = append(group[key], str)
	}

	var result = make([][]string, 0, len(group))
	for _, g := range group {
		result = append(result, g)
	}

	return result
}
