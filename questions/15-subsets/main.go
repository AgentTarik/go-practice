/*
Subsets

Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the subsets in any order.

Example 1:
  Input:  nums = [1, 2, 3]
  Output: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]

Example 2:
  Input:  nums = [0]
  Output: [[], [0]]

Constraints:
  - 1 <= nums.length <= 10
  - -10 <= nums[i] <= 10
  - All elements in nums are unique

Asked by: Amazon, Facebook, Bloomberg, Google
*/
package main

import (
	"fmt"
	"practice/testutil"
	"sort"
)

func subsets(nums []int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	type testCase struct {
		nums     []int
		expected [][]int
	}

	cases := []testCase{
		{[]int{1, 2, 3}, [][]int{
			{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
		}},
		{[]int{0}, [][]int{
			{}, {0},
		}},
		{[]int{1, 2}, [][]int{
			{}, {1}, {2}, {1, 2},
		}},
	}

	for _, tc := range cases {
		got := subsets(tc.nums)
		// sort both to compare regardless of order
		sortSets := func(sets [][]int) {
			for _, s := range sets {
				sort.Ints(s)
			}
			sort.Slice(sets, func(i, j int) bool {
				if len(sets[i]) != len(sets[j]) {
					return len(sets[i]) < len(sets[j])
				}
				for k := 0; k < len(sets[i]); k++ {
					if sets[i][k] != sets[j][k] {
						return sets[i][k] < sets[j][k]
					}
				}
				return false
			})
		}
		sortSets(got)
		sortSets(tc.expected)
		testutil.RunDeep(
			fmt.Sprintf("subsets(%v)", tc.nums),
			tc.expected,
			got,
		)
	}
}
