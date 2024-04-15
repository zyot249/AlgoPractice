package practice

/*
Problem:
Ref: https://leetcode.com/problems/merge-k-sorted-lists/description/

Given:
- An array of k linked lists
- Each list is sorted in ascending order

Return:
- Merge all to 1 and return it

Constraints:
- k: [0, 10^4]
- length of list: [0, 500]
- value of element in list: [-10^4, 10^4]
- sum of length of lists will not exceed 10^4
*/

/*
First approach:
I will iterate all the heads of lists, find the smallest node and add it to response list until reaching end of all input lists

Time complexity: O(mk^2)
Space complexity: O(1)
*/

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := &ListNode{}
	cur := result
	for {
		minIndex := 0
		for i := 1; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}

			if lists[minIndex] == nil || lists[i].Val < lists[minIndex].Val {
				minIndex = i
			}
		}

		if lists[minIndex] == nil {
			break
		}

		cur.Next = &ListNode{Val: lists[minIndex].Val}
		cur = cur.Next
		lists[minIndex] = lists[minIndex].Next
	}

	return result.Next
}

/*
Second approach:
Merge each list to result until reaching end

Improvement:
Instead of merging each list to result, we merge each 2 lists together.
Re-do this until all lists are merged to 1 list

Time complexity: O(mklogk)
Space complexity: O(1)
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	//result := lists[0]
	//for i := 1; i < len(lists); i++ {
	//	cur := result
	//	if lists[i] != nil {
	//		result = mergeTwoLists(cur, lists[i])
	//	}
	//}
	//
	//return result

	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], nil))
			}
		}

		lists = mergedLists
	}

	return lists[0]
}
