/*
Permutations

Given an array nums of distinct integers, return all possible permutations.
You can return the answer in any order.

Example 1:
  Input:  nums = [1, 2, 3]
  Output: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]

Example 2:
  Input:  nums = [0, 1]
  Output: [[0,1], [1,0]]

Example 3:
  Input:  nums = [1]
  Output: [[1]]

Constraints:
  - 1 <= nums.length <= 6
  - -10 <= nums[i] <= 10
  - All integers in nums are unique

Asked by: Amazon, Microsoft, Facebook, Bloomberg
*/
package main

import (
	"fmt"
	"practice/testutil"
	"sort"
)

func permute(nums []int) [][]int {
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
			{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
		}},
		{[]int{0, 1}, [][]int{
			{0, 1}, {1, 0},
		}},
		{[]int{1}, [][]int{
			{1},
		}},
	}

	for _, tc := range cases {
		got := permute(tc.nums)
		// sort both to compare regardless of order
		sortPerms := func(perms [][]int) {
			sort.Slice(perms, func(i, j int) bool {
				for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
					if perms[i][k] != perms[j][k] {
						return perms[i][k] < perms[j][k]
					}
				}
				return len(perms[i]) < len(perms[j])
			})
		}
		sortPerms(got)
		sortPerms(tc.expected)
		testutil.RunDeep(
			fmt.Sprintf("permute(%v)", tc.nums),
			tc.expected,
			got,
		)
	}
}
