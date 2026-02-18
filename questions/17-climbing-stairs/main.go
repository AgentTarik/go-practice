/*
Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Example 1:
  Input:  n = 2
  Output: 2  (1+1 or 2)

Example 2:
  Input:  n = 3
  Output: 3  (1+1+1, 1+2, 2+1)

Example 3:
  Input:  n = 5
  Output: 8

Constraints:
  - 1 <= n <= 45

Asked by: Amazon, Google, Apple, Adobe
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func climbStairs(n int) int {
	// TODO: implement
	return 0
}

func main() {
	type testCase struct {
		n        int
		expected int
	}

	cases := []testCase{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{10, 89},
		{20, 10946},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("climbStairs(%d)", tc.n),
			tc.expected,
			climbStairs(tc.n),
		)
	}
}
