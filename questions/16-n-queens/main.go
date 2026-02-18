/*
N-Queens

The n-queens puzzle is the problem of placing n queens on an n x n chessboard
such that no two queens attack each other. Given an integer n, return the number
of distinct solutions to the n-queens puzzle.

Example 1:
  Input:  n = 4
  Output: 2

Example 2:
  Input:  n = 1
  Output: 1

Example 3:
  Input:  n = 8
  Output: 92

Constraints:
  - 1 <= n <= 9

Asked by: Amazon, Google, Microsoft, Apple
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func totalNQueens(n int) int {
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
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{8, 92},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("totalNQueens(%d)", tc.n),
			tc.expected,
			totalNQueens(tc.n),
		)
	}
}
