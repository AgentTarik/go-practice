/*
Top K Frequent Elements (Heap / Priority Queue)

Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.

Example 1:
  Input:  nums = [1, 1, 1, 2, 2, 3], k = 2
  Output: [1, 2]

Example 2:
  Input:  nums = [1], k = 1
  Output: [1]

Example 3:
  Input:  nums = [4, 4, 4, 1, 1, 2, 2, 2, 3], k = 2
  Output: [4, 2]

Constraints:
  - 1 <= nums.length <= 10^5
  - -10^4 <= nums[i] <= 10^4
  - k is in the range [1, number of unique elements]
  - The answer is guaranteed to be unique

Asked by: Amazon, Facebook, Google, Apple
*/
package main

import (
	"fmt"
	"practice/testutil"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	// TODO: implement
	return nil
}

func main() {
	type testCase struct {
		nums     []int
		k        int
		expected []int
	}

	cases := []testCase{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 4, 4, 1, 1, 2, 2, 2, 3}, 2, []int{2, 4}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{5, 5, 5, 5, 1, 1, 1, 2, 2, 3}, 1, []int{5}},
	}

	for _, tc := range cases {
		got := topKFrequent(tc.nums, tc.k)
		sort.Ints(got)
		sort.Ints(tc.expected)
		testutil.RunDeep(
			fmt.Sprintf("topKFrequent(%v, %d)", tc.nums, tc.k),
			tc.expected,
			got,
		)
	}
}
