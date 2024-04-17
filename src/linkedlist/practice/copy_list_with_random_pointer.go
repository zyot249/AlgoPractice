package practice

import . "AlgoPractice/src/linkedlist/structure"

/*
Problem:
Ref: https://leetcode.com/problems/copy-list-with-random-pointer/description/

Given:
- A linked list
- Each node stores a random pointer to another node in the list or null

Return:
- The head of a deep copied list

Constraints:
- number of nodes: [0, 1000]
- node val: [-10^4, 10^4]
*/

/*
First approach:
I think I will use the hashmap to store the position of all nodes in original list.
When copying the node of original list to result, I also put the copied node to another array.
This will help me find the node that random pointer points to.

Time complexity: O(n)
Space complexity: O(n)

Improvement 1:
Instead of using a map and an arr, we use only 1 map to point each node in original node to the copy of it in response list
=> reduce using mem

Improvement 2:
We will use original list as the map from old node to copied node directly:
- Next of original node will point to its copy
- Next of its copy will point to next of it

Space complexity: O(1)
*/

func copyRandomList(head *Node) *Node {
	//res := &Node{}

	//nodePos := make(map[*Node]int)
	//var copiedArr []*Node
	//
	//ptr1 := head
	//ptr2 := res
	//nodeCount := 0
	//for ptr1 != nil {
	//	nodePos[ptr1] = nodeCount
	//	ptr2.Next = &Node{Val: ptr1.Val}
	//	copiedArr = append(copiedArr, ptr2.Next)
	//	ptr1 = ptr1.Next
	//	ptr2 = ptr2.Next
	//	nodeCount++
	//}
	//
	//ptr1 = head
	//ptr2 = res.Next
	//for ptr1 != nil {
	//	if ptr1.Random == nil {
	//		ptr2.Random = nil
	//	} else {
	//		index := nodePos[ptr1.Random]
	//		ptr2.Random = copiedArr[index]
	//	}
	//
	//	ptr1 = ptr1.Next
	//	ptr2 = ptr2.Next
	//}

	// Improvement 1
	//oldToCopy := make(map[*Node]*Node)
	//
	//ptr1 := head
	//ptr2 := res
	//for ptr1 != nil {
	//	ptr2.Next = &Node{Val: ptr1.Val}
	//	oldToCopy[ptr1] = ptr2.Next
	//	ptr1 = ptr1.Next
	//	ptr2 = ptr2.Next
	//}
	//
	//ptr1 = head
	//for ptr1 != nil {
	//	copyNode := oldToCopy[ptr1]
	//	if ptr1.Random == nil {
	//		copyNode.Random = nil
	//	} else {
	//		copyNode.Random = oldToCopy[ptr1.Random]
	//	}
	//
	//	ptr1 = ptr1.Next
	//}

	//Improvement 2
	if head == nil {
		return nil
	}

	ptr1 := head
	for ptr1 != nil {
		copyNode := &Node{Val: ptr1.Val, Next: ptr1.Next}
		ptr1.Next = copyNode
		ptr1 = copyNode.Next
	}

	ptr1 = head
	for ptr1 != nil {
		copyNode := ptr1.Next
		if ptr1.Random != nil {
			copyNode.Random = ptr1.Random.Next
		}

		ptr1 = copyNode.Next
	}

	ptr1 = head
	res := head.Next
	for ptr1 != nil {
		copyNode := ptr1.Next
		nextNode := copyNode.Next
		if nextNode != nil {
			copyNode.Next = nextNode.Next
		}

		ptr1.Next = nextNode
		ptr1 = nextNode
	}

	return res
}
