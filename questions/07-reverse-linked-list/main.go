/*
Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

	Input:  [1, 2, 3, 4, 5]
	Output: [5, 4, 3, 2, 1]

Example 2:

	Input:  [1, 2]
	Output: [2, 1]

Example 3:

	Input:  []
	Output: []

Constraints:
  - The number of nodes in the list is in the range [0, 5000]
  - -5000 <= Node.val <= 5000

Asked by: Amazon, Microsoft, Apple, Bloomberg
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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextNode := head.Next
	head.Next = nil

	for nextNode.Next != nil {
		temp1 := nextNode.Next
		temp2 := nextNode
		nextNode.Next = head
		head = temp2
		nextNode = temp1
	}
	nextNode.Next = head
	return nextNode
}

// helpers

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	type testCase struct {
		input    []int
		expected []int
	}

	cases := []testCase{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int(nil)},
	}

	for _, tc := range cases {
		testutil.RunDeep(
			fmt.Sprintf("reverseList(%v)", tc.input),
			tc.expected,
			listToSlice(reverseList(buildList(tc.input))),
		)
	}
}
