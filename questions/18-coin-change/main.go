/*
Coin Change

You are given an integer array coins representing coin denominations and an integer
amount representing a total amount of money. Return the fewest number of coins
needed to make up that amount. If that amount cannot be made up by any combination
of the coins, return -1. You may use each coin an unlimited number of times.

Example 1:
  Input:  coins = [1, 5, 10, 25], amount = 30
  Output: 2  (25 + 5)

Example 2:
  Input:  coins = [2], amount = 3
  Output: -1

Example 3:
  Input:  coins = [1], amount = 0
  Output: 0

Constraints:
  - 1 <= coins.length <= 12
  - 1 <= coins[i] <= 2^31 - 1
  - 0 <= amount <= 10^4

Asked by: Amazon, Google, Microsoft, Apple
*/
package main

import (
	"fmt"
	"practice/testutil"
)

func coinChange(coins []int, amount int) int {
	// TODO: implement
	return -1
}

func main() {
	type testCase struct {
		coins    []int
		amount   int
		expected int
	}

	cases := []testCase{
		{[]int{1, 5, 10, 25}, 30, 2},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 11, 3},         // 5+5+1
		{[]int{2}, 4, 2},                 // 2+2
		{[]int{1}, 1, 1},
		{[]int{1, 2, 5}, 100, 20},        // 20 x 5
		{[]int{186, 419, 83, 408}, 6249, 20},
	}

	for _, tc := range cases {
		testutil.Run(
			fmt.Sprintf("coinChange(%v, %d)", tc.coins, tc.amount),
			tc.expected,
			coinChange(tc.coins, tc.amount),
		)
	}
}
