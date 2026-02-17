/*
Two Sum

Given an array of integers nums and an integer target, return the indices of the
two numbers that add up to target. You may assume exactly one solution exists,
and you may not use the same element twice.

Example 1:

	Input:  nums = [2, 7, 11, 15], target = 9
	Output: [0, 1]  (nums[0] + nums[1] == 9)

Example 2:

	Input:  nums = [3, 2, 4], target = 6
	Output: [1, 2]

Example 3:

	Input:  nums = [3, 3], target = 6
	Output: [0, 1]

Constraints:
  - 2 <= nums.length <= 10^4
  - -10^9 <= nums[i] <= 10^9
  - -10^9 <= target <= 10^9
  - Exactly one valid answer exists
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func twoSum(nums []int, target int) [2]int {
	numMap := make(map[int]int)
	for i, num := range nums {
		lookup := target - num
		value, ok := numMap[lookup]
		if ok && i != value {
			return [2]int{i, value}
		}
		numMap[num] = i
	}

	return [2]int{-1, -1}
}

func main() {
	type testCase struct {
		nums     []int
		target   int
		expected [2]int
	}

	cases := []testCase{
		{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},        // basic
		{[]int{3, 2, 4}, 6, [2]int{1, 2}},             // answer not at start
		{[]int{3, 3}, 6, [2]int{0, 1}},                // duplicate values
		{[]int{-3, 4, 3, 90}, 0, [2]int{0, 2}},        // negative + positive sum to zero
		{[]int{-1, -2, -3, -4, -5}, -8, [2]int{2, 4}}, // all negatives
		{[]int{0, 4, 3, 0}, 0, [2]int{0, 3}},          // zeros
		{[]int{1, 2, 3, 4, 5}, 9, [2]int{3, 4}},       // answer at the end
	}

	normalize := func(p [2]int) [2]int {
		if p[0] > p[1] {
			return [2]int{p[1], p[0]}
		}
		return p
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("twoSum(%v, %d)", tc.nums, tc.target),
			normalize(tc.expected),
			normalize(twoSum(tc.nums, tc.target)),
		)
	}
}
