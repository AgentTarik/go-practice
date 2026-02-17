/*
Binary Search

Given a sorted array of integers and a target value, return the index of the target.
If the target is not found, return -1. Must run in O(log n) time.

Example 1:
  Input:  nums = [-1, 0, 3, 5, 9, 12], target = 9
  Output: 4

Example 2:
  Input:  nums = [-1, 0, 3, 5, 9, 12], target = 2
  Output: -1

Example 3:
  Input:  nums = [5], target = 5
  Output: 0

Constraints:
  - 1 <= nums.length <= 10^4
  - -10^4 <= nums[i], target <= 10^4
  - All values in nums are unique
  - nums is sorted in ascending order
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func binarySearch(nums []int, target int) int {
	// TODO: implement
	return -1
}

func main() {
	type testCase struct {
		nums     []int
		target   int
		expected int
	}

	cases := []testCase{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},   // found in middle
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},  // not found
		{[]int{5}, 5, 0},                     // single element: found
		{[]int{5}, 6, -1},                    // single element: not found
		{[]int{1, 3, 5, 7, 9, 11}, 1, 0},    // leftmost element
		{[]int{1, 3, 5, 7, 9, 11}, 11, 5},   // rightmost element
		{[]int{1, 3, 5, 7, 9, 11}, 0, -1},   // less than all elements
		{[]int{1, 3, 5, 7, 9, 11}, 12, -1},  // greater than all elements
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("binarySearch(%v, %d)", tc.nums, tc.target),
			tc.expected,
			binarySearch(tc.nums, tc.target),
		)
	}
}
