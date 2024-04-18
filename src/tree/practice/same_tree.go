package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/same-tree/description/

Given:
- The roots of 2 binary trees p and q
- Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Return:
- True if they are the same and false otherwise

Constraints:
- Number of nodes: [0, 100]
- Node value: [-10^4, 10^4]
*/

/*
First approach:
I will traverse 2 trees at the same time with the same traversal type
2 trees are the same if they are the same at every position

Time complexity: O(n)
Space complexity: O(1)
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	} else if q != nil && p != nil {
		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
