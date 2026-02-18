/*
Merge K Sorted Lists (Heap / Priority Queue)

You are given an array of k linked lists, each sorted in ascending order.
Merge all the linked lists into one sorted linked list and return it.

Example 1:
  Input:  lists = [[1,4,5], [1,3,4], [2,6]]
  Output: [1, 1, 2, 3, 4, 4, 5, 6]

Example 2:
  Input:  lists = []
  Output: []

Example 3:
  Input:  lists = [[]]
  Output: []

Constraints:
  - 0 <= k <= 10^4
  - 0 <= lists[i].length <= 500
  - -10^4 <= lists[i][j] <= 10^4
  - lists[i] is sorted in ascending order
  - The sum of lists[i].length will not exceed 10^4

Asked by: Amazon, Facebook, Google, Microsoft
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

func mergeKLists(lists []*ListNode) *ListNode {
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
		lists    [][]int
		expected []int
	}

	cases := []testCase{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, nil},
		{[][]int{{}}, nil},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{-1, 0, 3}, {-2, 5}}, []int{-2, -1, 0, 3, 5}},
	}

	for _, tc := range cases {
		var lists []*ListNode
		for _, vals := range tc.lists {
			lists = append(lists, buildList(vals))
		}
		testutil.RunDeep(
			fmt.Sprintf("mergeKLists(%v)", tc.lists),
			tc.expected,
			listToSlice(mergeKLists(lists)),
		)
	}
}
