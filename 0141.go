package lc0141

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		vec := make([]*TreeNode, 0, 64)
		vec = prevIter(vec, root)

		var node, prev *TreeNode
		for i := 0; i <= len(vec)-1; i++ {
			node = vec[i]
			(*node).Left = nil

			if prev != nil {
				(*prev).Right = node
			}

			prev = node
		}
		(*node).Right = nil
	}
}

func prevIter(vec []*TreeNode, root *TreeNode) []*TreeNode {
	if root != nil {
		vec = append(vec, root)
		vec = prevIter(vec, root.Left)
		vec = prevIter(vec, root.Right)
	}
	return vec
}

func flatten2(root *TreeNode) {
	list := []*TreeNode{}
	stack := []*TreeNode{}
	node := root
	// $$ 递归转迭代
	for node != nil || len(stack) > 0 {
		// $$ 处理这个node
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}

		// $$ node 重新赋值，利用栈顶信息
		node = stack[len(stack)-1].Right

		// $$ 栈顶元素弹出，待处理
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

// $$ 递归转迭代，中序版
func middleTraverse(r *TreeNode) {
	s := []*TreeNode{}
	node := r
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			// if node.Left == nil {
			// 	// do sth to node
			// 	fmt.Println(node)
			// 	break
			// } else {
			// 	node = node.Left
			// }
			node = node.Left
		}

		// pop from vec, deal with the right child of the poped element
		if len(s) > 0 {
			// do sth to top element
			fmt.Println(s[len(s)-1])

			node = s[len(s)-1].Right

			s = s[:len(s)-1]
		}
	}
}

func flatten3(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			prev := next
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
func flatten4(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// prev is the predecessor of curr.Right
			prev := curr.Left
			for prev.Right != nil {
				prev = prev.Right
			}

			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}

		curr = curr.Right
	}
}
