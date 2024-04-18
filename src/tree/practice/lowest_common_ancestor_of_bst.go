package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

Given:
- a BST
- node p and q
- The lowest common ancestor (LCA) is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)

Return:
- The LCA of p and q

Constraints:
- Number of nodes: [2, 10^5]
- All node values are unique
- p != q
- p and q will exist in BST
*/

/*
First approach:
I will find all the ancestors of q and p
From that, I can find the LCA of them

Time complexity: O(2logn)
Space complexity: O(1)
*/

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p == nil {
		return q
	}

	if q == nil {
		return p
	}

	var result *TreeNode
	findP, findQ := root, root
	for findP != nil && findQ != nil {
		if findP == findQ {
			result = findP
		} else {
			break
		}

		if findP.Val < p.Val {
			findP = findP.Right
		} else if findP.Val > p.Val {
			findP = findP.Left
		}

		if findQ.Val < q.Val {
			findQ = findQ.Right
		} else if findQ.Val > q.Val {
			findQ = findQ.Left
		}
	}

	return result
}

/*
Solution:
If p and q < root -> LCA of p and q is in root.Left
If p and q > root -> LCA of p and q is in root.Right
Else LCA of q and p is root

Time complexity: O(logn)
Space complexity: O(1)
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
