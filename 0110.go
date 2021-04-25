package lc110

// #哨兵法改写
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	lbool := isBalanced(root.Left)
// 	rbool := isBalanced(root.Right)
// 	diff := depth(root.Left) - depth(root.Right)
// 	return (lbool && rbool) && (diff <= 1 && diff >= -1)

// }
// func depth(root *TreeNode) int {
// 	if root != nil {
// 		return max(depth(root.Left), depth(root.Right)) + 1
// 	}
// 	return 0
// }
// func max(a int, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root != nil {
		lbool := isBalanced(root.Left)
		if !lbool {
			return false
		}

		rbool := isBalanced(root.Right)
		if !rbool {
			return false
		}

		diff := depth(root.Left) - depth(root.Right)
		return (diff <= 1 && diff >= -1)
	}
	return true
}
func depth(root *TreeNode) int {
	if root != nil {
		return max(depth(root.Left), depth(root.Right)) + 1
	}
	return 0
}
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
