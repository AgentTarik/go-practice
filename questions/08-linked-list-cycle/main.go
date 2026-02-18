/*
Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.
A cycle exists if some node in the list can be reached again by continuously following
the next pointer. Return true if there is a cycle, false otherwise.

Example 1:
  Input:  [3, 2, 0, -4], cycle at position 1
  Output: true  (tail connects to node at index 1)

Example 2:
  Input:  [1, 2], cycle at position 0
  Output: true  (tail connects to node at index 0)

Example 3:
  Input:  [1], no cycle
  Output: false

Constraints:
  - The number of nodes is in the range [0, 10^4]
  - -10^5 <= Node.val <= 10^5
  - Can you solve it with O(1) memory?

Asked by: Amazon, Microsoft, Bloomberg, Goldman Sachs
*/
package main

import (
	"fmt"
	"practice/testutil"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// TODO: implement
	return false
}

// helpers

func buildCycleList(vals []int, cyclePos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if cyclePos >= 0 {
		nodes[len(nodes)-1].Next = nodes[cyclePos]
	}
	return nodes[0]
}

func main() {
	type testCase struct {
		vals     []int
		cyclePos int
		expected bool
	}

	cases := []testCase{
		{[]int{3, 2, 0, -4}, 1, true},  // tail connects to index 1
		{[]int{1, 2}, 0, true},          // tail connects to index 0
		{[]int{1}, -1, false},           // single node, no cycle
		{[]int{1, 2, 3, 4, 5}, -1, false}, // no cycle
		{[]int{1, 2, 3, 4, 5}, 4, true},   // tail connects to itself
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("hasCycle(%v, cyclePos=%d)", tc.vals, tc.cyclePos),
			tc.expected,
			hasCycle(buildCycleList(tc.vals, tc.cyclePos)),
		)
	}
}
