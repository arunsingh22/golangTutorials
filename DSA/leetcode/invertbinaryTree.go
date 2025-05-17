package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// swap the children of the root node
	tmpNode := root.Left
	root.Left = root.Right
	root.Right = tmpNode

	// call recursively for the children now
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func BFSInvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// this gives a doubly linked list.
	qu := []*TreeNode{}
	qu = append(qu, root)

	for len(qu) > 0 {
		frontNode := qu[0]
		qu = qu[1:]

		// swap it's children
		frontNode.Left, frontNode.Right = frontNode.Right, frontNode.Left

		if frontNode.Left != nil {
			qu = append(qu, frontNode.Left)
		}
		if frontNode.Right != nil {
			qu = append(qu, frontNode.Right)
		}

	}

	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// this gives a doubly linked list.
	qu := []*TreeNode{}
	qu = append(qu, root)
	cnt := 1
	for len(qu) > 0 {
		frontNode := qu[0]
		qu = qu[1:]

		if cnt&1 != 0 {
			// swap it's children
			lc := frontNode.Left
			rc := frontNode.Right
			lc.Val, rc.Val = rc.Val, lc.Val
			// frontNode.Left.Value, frontNode.Right.Value = frontNode.Right.Value, frontNode.Left.Value
		}

		if frontNode.Left != nil {
			qu = append(qu, frontNode.Left)
		}
		if frontNode.Right != nil {
			qu = append(qu, frontNode.Right)
		}
		cnt++
	}
	return root
}
