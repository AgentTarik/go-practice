/*
Maximum Depth of Binary Tree (DFS)

Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.

Example 1:
  Input:  root = [3, 9, 20, nil, nil, 15, 7]
  Output: 3

Example 2:
  Input:  root = [1, nil, 2]
  Output: 2

Example 3:
  Input:  root = []
  Output: 0

Constraints:
  - The number of nodes is in the range [0, 10^4]
  - -100 <= Node.val <= 100

Asked by: Amazon, Google, Microsoft, LinkedIn
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

func maxDepth(root *TreeNode) int {
	// TODO: implement
	return 0
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

func main() {
	type testCase struct {
		vals     []int
		expected int
	}

	null := -101

	cases := []testCase{
		{[]int{3, 9, 20, null, null, 15, 7}, 3},
		{[]int{1, null, 2}, 2},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5, null, null}, 3},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("maxDepth(%v)", tc.vals),
			tc.expected,
			maxDepth(buildTree(tc.vals, null)),
		)
	}
}
