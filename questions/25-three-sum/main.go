/*
3Sum (Two Pointers)

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, j != k, and nums[i] + nums[j] + nums[k] == 0.
The solution set must not contain duplicate triplets.

Example 1:
  Input:  nums = [-1, 0, 1, 2, -1, -4]
  Output: [[-1, -1, 2], [-1, 0, 1]]

Example 2:
  Input:  nums = [0, 1, 1]
  Output: []

Example 3:
  Input:  nums = [0, 0, 0]
  Output: [[0, 0, 0]]

Constraints:
  - 3 <= nums.length <= 3000
  - -10^5 <= nums[i] <= 10^5

Asked by: Amazon, Facebook, Google, Microsoft
*/
package main

import (
	"fmt"
	"practice/testutil"
	"sort"
)

func threeSum(nums []int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	type testCase struct {
		nums     []int
		expected [][]int
	}

	cases := []testCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, nil},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{[]int{1, 2, -2, -1}, nil},
	}

	for _, tc := range cases {
		got := threeSum(tc.nums)
		// sort each triplet and the list for comparison
		sortTrips := func(trips [][]int) {
			for _, t := range trips {
				sort.Ints(t)
			}
			sort.Slice(trips, func(i, j int) bool {
				for k := 0; k < 3; k++ {
					if trips[i][k] != trips[j][k] {
						return trips[i][k] < trips[j][k]
					}
				}
				return false
			})
		}
		sortTrips(got)
		sortTrips(tc.expected)
		testutil.RunDeep(
			fmt.Sprintf("threeSum(%v)", tc.nums),
			tc.expected,
			got,
		)
	}
}
