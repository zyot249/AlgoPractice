package practice

/*
Problem:
Ref: https://leetcode.com/problems/longest-consecutive-sequence/
Given:
- An unsorted array of integers

Return:
- The length of the longest consecutive sequence

Constraints:
- arr length: [0, 10^5]
- arr[i]: [-10^9, 10^9]
- Time complexity: O(n)
*/

/*
First approach:
Firstly, we convert the array into a set of numbers for checking the existence
Secondly, we iterate all the elements in array.
For each element, if it isn't a start point of sequence, skip it.
If it is a start point of sequence, find next element in its sequence until reaching end and calculate the length of sequence

Time complexity: O(n)
Space complexity: O(n)
*/
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	maxLength := 0
	//Solution 1
	//for num, _ := range set {
	//	if set[num-1] {
	//		continue
	//	}
	//
	//	length := 1
	//	for next := num + 1; set[next] == true; next++ {
	//		length++
	//	}
	//
	//	if length > maxLength {
	//		maxLength = length
	//	}
	//}

	//Solution 2
	for _, num := range nums {
		if !set[num] {
			continue
		}

		delete(set, num)
		length := 1

		for next := num + 1; set[next] == true; next++ {
			delete(set, next)
			length++
		}

		for prev := num - 1; set[prev] == true; prev-- {
			delete(set, prev)
			length++
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
