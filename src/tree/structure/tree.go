package structure

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Integer struct {
	Value int
}

func CreateBinaryTree(vals []*Integer) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := createBinarySubTree(vals, 0)
	return root
}

func createBinarySubTree(vals []*Integer, i int) *TreeNode {
	if i >= len(vals) || vals[i] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[i].Value}
	root.Left = createBinarySubTree(vals, 2*i+1)
	root.Right = createBinarySubTree(vals, 2*i+2)
	return root
}
