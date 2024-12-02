package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return dfs(root.Left, val)
	}
	return dfs(root.Right, val)
}

func searchBST(root *TreeNode, val int) *TreeNode {
	return dfs(root, val)
}
