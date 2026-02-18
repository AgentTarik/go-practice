/*
Invert Binary Tree

Given the root of a binary tree, invert the tree and return its root.
Inverting means swapping every left child with its right child.

Example 1:
  Input:  root = [4, 2, 7, 1, 3, 6, 9]
  Output: [4, 7, 2, 9, 6, 3, 1]

Example 2:
  Input:  root = [2, 1, 3]
  Output: [2, 3, 1]

Example 3:
  Input:  root = []
  Output: []

Constraints:
  - The number of nodes is in the range [0, 100]
  - -100 <= Node.val <= 100

Asked by: Google, Amazon, Facebook, Apple
*/
package main

import (
	"fmt"
	"practice/testutil"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// TODO: implement
	return nil
}

// helpers

func buildTree(vals []int, nilVal int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nilVal {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nilVal {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func main() {
	type testCase struct {
		vals     []int
		expected []int
	}

	null := -101

	cases := []testCase{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{}, nil},
		{[]int{1}, []int{1}},
	}

	for _, tc := range cases {
		testutil.RunDeep(
			fmt.Sprintf("invertTree(%v)", tc.vals),
			tc.expected,
			treeToSlice(invertTree(buildTree(tc.vals, null))),
		)
	}
}
