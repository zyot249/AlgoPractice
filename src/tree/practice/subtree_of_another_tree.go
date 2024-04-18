package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/subtree-of-another-tree/description/

Given:
- The roots of 2 trees root and subRoot

Return:
- true if there is a subtree of root with the same structure and node values of subRoot and false otherwise

Constraints:
- number of nodes in root: [1, 2000]
- number of nodes in subRoot: [1, 1000]
- node value: [-10^4, 10^4]
*/

/*
First approach:
I will traverse the root tree and compare its nodes to the subRoot

Time complexity: O(m.n)
Space complexity: O(1)
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) ||
		(root.Left != nil && isSubtree(root.Left, subRoot)) ||
		(root.Right != nil && isSubtree(root.Right, subRoot)) {
		return true
	}

	return false
}
