/*
Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by
splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
  Input:  list1 = [1, 2, 4], list2 = [1, 3, 4]
  Output: [1, 1, 2, 3, 4, 4]

Example 2:
  Input:  list1 = [], list2 = []
  Output: []

Example 3:
  Input:  list1 = [], list2 = [0]
  Output: [0]

Constraints:
  - The number of nodes in both lists is in the range [0, 50]
  - -100 <= Node.val <= 100
  - Both list1 and list2 are sorted in non-decreasing order

Asked by: Amazon, Microsoft, Apple, Facebook
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: implement
	return nil
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
		list1    []int
		list2    []int
		expected []int
	}

	cases := []testCase{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int(nil)},
		{[]int{}, []int{0}, []int{0}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{5, 10, 15}, []int{1, 2, 3}, []int{1, 2, 3, 5, 10, 15}},
	}

	for _, tc := range cases {
		testutil.RunDeep(
			fmt.Sprintf("mergeTwoLists(%v, %v)", tc.list1, tc.list2),
			tc.expected,
			listToSlice(mergeTwoLists(buildList(tc.list1), buildList(tc.list2))),
		)
	}
}
