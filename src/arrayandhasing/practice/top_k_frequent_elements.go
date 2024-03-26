package practice

import (
	"sort"
)

/*
Problem:
Ref: https://leetcode.com/problems/top-k-frequent-elements/
Given
- integer arr
- integer k

-> return k most frequent elements of arr

Constraints:
- arr length: [0, 10^5]
- arr[i]: [-10^4, 10^4]
- k: [1, num of unique elements in arr]
- result is unique
*/

/*
First approach:
We store the frequency of all elements in arr in a Hashmap
After that, we sort the list of pairs in hashmap by value
Get top k results

Time complexity: O(n + nlogn + k) ~ O(nlogn)
Space complexity: O(n)
*/

type Pair struct {
	Key   int
	Value int
}

func topKFrequent1(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	var freArr []Pair
	for k, v := range fre {
		freArr = append(freArr, Pair{k, v})
	}

	sort.Slice(freArr, func(i int, j int) bool {
		a := freArr[i]
		b := freArr[j]
		return a.Value > b.Value
	})

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, freArr[i].Key)
	}

	return result
}

/*
Second approach:
Using bucket sort
In order to find top frequent elements without sorting, we can convert the map of number and its frequency into an array
by set the value of element at index frequency is the numbers that have corresponding frequency.
After that, we iterate from the end of arr for finding top k

Time complexity: O(n)
Space complexity: O(n)
*/
func topKFrequent(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, f := range fre {
		buckets[f] = append(buckets[f], num)
	}

	var result []int
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(result) >= k {
			return result
		}

		if len(buckets[i]) > 0 {
			for _, num := range buckets[i] {
				result = append(result, num)
			}
		}
	}

	return result
}
