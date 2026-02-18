/*
Lowest Common Ancestor of a Binary Tree

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes.
The LCA is the deepest node that is an ancestor of both p and q
(a node can be an ancestor of itself).

Example 1:
  Input:  root = [3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4], p = 5, q = 1
  Output: 3

Example 2:
  Input:  root = [3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4], p = 5, q = 4
  Output: 5  (5 is an ancestor of itself)

Example 3:
  Input:  root = [1, 2], p = 1, q = 2
  Output: 1

Constraints:
  - The number of nodes is in the range [2, 10^5]
  - -10^9 <= Node.val <= 10^9
  - All Node.val are unique
  - p != q
  - p and q will exist in the tree

Asked by: Facebook, Amazon, Microsoft, Google
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	type testCase struct {
		vals     []int
		pVal     int
		qVal     int
		expected int
	}

	null := -1000000001

	cases := []testCase{
		{[]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}, 5, 1, 3},
		{[]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}, 5, 4, 5},
		{[]int{1, 2}, 1, 2, 1},
		{[]int{3, 5, 1}, 5, 1, 3},
	}

	for _, tc := range cases {
		root := buildTree(tc.vals, null)
		p := findNode(root, tc.pVal)
		q := findNode(root, tc.qVal)
		result := lowestCommonAncestor(root, p, q)
		resultVal := -1
		if result != nil {
			resultVal = result.Val
		}
		testutil.Run(
			fmt.Sprintf("LCA(%v, p=%d, q=%d)", tc.vals, tc.pVal, tc.qVal),
			tc.expected,
			resultVal,
		)
	}
}
