package main

import (
	"fmt"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {
	// preorder traversal

	ans := [][]int{}
	keys := []int{}
	if root == nil {
		return ans
	}

	vl := 0
	dict := map[int][]int{} //composite literal
	keys = append(keys, vl)
	dict[vl] = []int{root.Val}

	dfs(root.Left, vl-1, dict, &keys)
	dfs(root.Right, vl+1, dict, &keys)

	fmt.Println(dict)

	// sort keys and and it's values
	slices.Sort(keys)
	// fmt.Println("keys: ", keys)
	for i := 0; i < len(keys); i++ {
		val, ok := dict[keys[i]]
		if ok {
			slices.Sort(val)
			ans = append(ans, val)
		}
		// fmt.Println(ans)
	}
	return ans
}

func dfs(root *TreeNode, vl int, dict map[int][]int, keys *[]int) {
	if root == nil {
		return
	}
	if val, ok := dict[vl]; ok {
		// value is present
		dict[vl] = append(val, root.Val)
	} else {
		dict[vl] = []int{root.Val}
		*keys = append(*keys, vl)
	}
	dfs(root.Left, vl-1, dict, keys)
	dfs(root.Right, vl+1, dict, keys)
}
