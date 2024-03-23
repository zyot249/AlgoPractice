package practice

import (
	"math"
	"sort"
)

/*
First approach:
Iterate all the elements of arr
For each value of element, I will check if I met it before by using an arr
Since the value of element in range of -10^9 -> 10^9 --> size of checkExistArr should be 2*10^9
I will use the value of element as index of checkExistArr
The time complexity will be O(n)
The memory complexity will be O(2m)
*/
func containsDuplicate1(nums []int) bool {
	var arrLen = int(2 * math.Pow(10, 9))
	var checkExistArr = make([]bool, arrLen)
	for i := 0; i < arrLen; i++ {
		checkExistArr[i] = false
	}

	for _, num := range nums {
		var index = num + int(math.Pow(10, 9))
		if checkExistArr[index] == true {
			return true
		}

		checkExistArr[index] = true
	}

	return false
}

/*
Second approach:
Iterate all the elements of arr
For each element in arr, I will iterate from its index to the end of arr to check if having any element that has the same value with it
The time complexity will be O(n^2)
The memory complexity will be O(1)
*/

func containsDuplicate2(nums []int) bool {
	for index, num := range nums {
		for i := index + 1; i < len(nums); i++ {
			if nums[i] == num {
				return true
			}
		}
	}

	return false
}

/*
Third approach:
Iterate all the elements of arr
For each value of element, I will check if I met it before by using a hashmap
Since the value of element in range of -10^9 -> 10^9 --> size of checkExistArr should be 2*10^9
I will use the value of element as index of checkExistArr
The time complexity will be O(n)
The memory complexity will be O(n)
*/
func containsDuplicate3(nums []int) bool {
	var lookup = make(map[int]bool, len(nums))
	for _, num := range nums {
		if lookup[num] {
			return true
		}

		lookup[num] = true
	}

	return false
}

/*
Forth approach:
Firstly, I will sort the arr
After that, I iterate all the elements of arr.
If element at i + 1 == element at i ==> return true
The time complexity will be O(nlogn)
The memory complexity will be O(1)
*/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
