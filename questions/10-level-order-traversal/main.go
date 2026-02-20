/*
Binary Tree Level Order Traversal (BFS)

Given the root of a binary tree, return the level order traversal of its
nodes' values (i.e., from left to right, level by level).

Example 1:

	Input:  root = [3, 9, 20, nil, nil, 15, 7]
	Output: [[3], [9, 20], [15, 7]]

Example 2:

	Input:  root = [1]
	Output: [[1]]

Example 3:

	Input:  root = []
	Output: []

Constraints:
  - The number of nodes is in the range [0, 2000]
  - -1000 <= Node.val <= 1000

Asked by: Amazon, Microsoft, Facebook, Bloomberg
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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		innerResult := []int{}
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			innerResult = append(innerResult, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, innerResult)
	}
	return result
}

// helpers

// buildTree builds a binary tree from a level-order slice.
// Use -1 as a placeholder for nil nodes.
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
		expected [][]int
	}

	null := -101 // sentinel for nil nodes

	cases := []testCase{
		{[]int{3, 9, 20, null, null, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{}, nil},
		{[]int{1, 2, 3, 4, 5}, [][]int{{1}, {2, 3}, {4, 5}}},
	}

	for _, tc := range cases {
		testutil.RunDeep(
			fmt.Sprintf("levelOrder(%v)", tc.vals),
			tc.expected,
			levelOrder(buildTree(tc.vals, null)),
		)
	}
}
